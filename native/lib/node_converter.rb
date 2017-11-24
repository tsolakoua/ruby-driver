require 'parser/current'

# Nodes doc:
# https://github.com/whitequark/parser/blob/master/doc/AST_FORMAT.md

module NodeConverter
  class Converter
    def initialize(node)
      @root = node
      @dict = {}
      if not @root.is_a?(Parser::AST::Node)
        raise "Object is not a Parser::AST::Node"
      end
    end

    def tohash()
      @dict["root"] = convert(@root)
    end

    private

    def convert(node)
      type = node_type(node)

      case type

      when "int", "float", "str"
        return sexp_to_hash(node, {"l_token" => 0})

      when "lvar", "ivar", "cvar", "gvar", "arg", "kwarg", "restarg", "blockarg"
        return sexp_to_hash(node, {"token.token" => 0}, 1, "children")

      when "pair", "irange", "erange", "alias", "iflipflop", "eflipflop"
        return sexp_to_hash(node, {"_1" => 0, "_2" => 1})

      when "lvasgn", "ivasgn", "cvasgn"
        return sexp_to_hash(node, {"target" => 0, "value" => 1})

      when "or_asgn", "and_asgn"
        return sexp_to_hash(node, {"target.target" => 0, "value" => 1})

      when "optarg", "kwoptarg"
        return sexp_to_hash(node, {"token" => 0, "default" => 1})

      when "splat", "kwsplat", "defined?", "kwrestarg"
        return sexp_to_hash(node, {"name" => 0})

      when "casgn", "csend", "send"
        return sexp_to_hash(node, {"base" => 0, "selector" => 1, "value" => 2})

      when "complex", "rational", "sym"
        return sexp_to_hash(node, {"token.token" => 0})

      # the inner nodes of the above
      when "Complex", "Rational", "Symbol"
        return {"type" => node_type(node), "token" => node.to_s}

      when "mlhs"
        return node.children.map{ |x| convert(x) }.compact

      when "masgn"
        return sexp_to_hash(node, {"targets" => 0, "values" => 1})

      when "op_asgn"
        return sexp_to_hash(node, {"target.target" => 0, "operator" => 1, "value" => 2})

      when "module"
        d = sexp_to_hash(node, {}, 1, "body")
        d["name"] = node.children[0].children[1].to_s
        return d

      when "class"
        d = sexp_to_hash(node, {"parent" => 1}, 2, "body")
        d["name"] = node.children[0].children[1].to_s
        return d

      when "sclass"
        return sexp_to_hash(node, {"object" => 0}, 1, "body")

      when "def"
        return sexp_to_hash(node, {"s_name" => 0, "args" => 1}, 2, "body")

      when "undef", "yield", "break", "next", "return"
        return sexp_to_hash(node, {"target" => 0})

      when "and", "or"
        return sexp_to_hash(node, {"left" => 0, "right" => 1})

      when "case"
        d = sexp_to_hash(node, {"casevar" => 0})
        if node.children.length > 2
          d["when_clauses"] = node.children[1..-2].map{ |x| convert(x) }.compact
        end
        d["else"] = convert(node.children[-1])
        return d

      when "const"
        return sexp_to_hash(node, {"base" => 0, "token" => 1})

      when "while", "until", "while_post", "until_post"
        return sexp_to_hash(node, {"condition" => 0, "body" => 1})

      when "begin", "kwbegin", "preexe", "postexe"
        return sexp_to_hash(node, {}, 0, "body")

      when "for"
        return sexp_to_hash(node, {"iterators" => 0, "iterated" => 1, "body" => 2})

      when "resbody"
        return sexp_to_hash(node, {"exceptions" => 0, "alias" => 1, "body" => 2})

      when "rescue"
        d = sexp_to_hash(node, {"body" => 0})
        if node.children.length > 2
          d["handlers"] = node.children[1..-2].map{ |x| convert(x) }.compact
        end
        d["else"] = convert(node.children[-1])
        return d

      when "ensure"
        return sexp_to_hash(node, {"body" => 0, "ensure_body" => 1})

      when "if"
        return sexp_to_hash(node, {"condition" => 0, "body" => 1, "else" => 2})

      else
        # default conversion
        return sexp_to_hash(node, {}, 0, "children")
      end
    end

    def node_type(node)
      (node.is_a?(Parser::AST::Node) ? node.type : node.class).to_s
    end

    # Convert positional children nodes to hashtable nodes keyed to a named
    # attribute using a attrname => position hashmap in the "table" argument.
    #
    # The attrnames can use a micro-DSL to specify further operations on the
    # children nodes:
    #
    # - Start with "s_": the children node wont be visited and will be converted
    #   to string as is.
    # - Starts with "l_": take child literal value, without calling convert.
    # - Ends with ".foo": the child node will be visited and the key "foo" from
    #   the resulting dictionary will be assigned.
    #
    # In any other case, the child will just be converted and assigned to the specified
    # key in the node dictionary.
    #
    # The cdr_index and cdr_key arguments, if present, specify that any other children
    # after cdr_index will be converted and assigned as a list of dictnodes to the cdr_key
    # property in the node.
    def sexp_to_hash(node, table, cdr_index=nil, cdr_key=nil)
      d = {"type" => node_type(node)}

      table.each do |propname, idx|
        if propname.start_with? "s_"
          d[propname[2..-1]] = node.children[idx].to_s

        elsif propname.start_with? "l_"
          d[propname[2..-1]] = node.children[idx]

        elsif propname.include? "."
          propname, childkey = propname.split(".")
          d[propname] = convert(node.children[idx])[childkey]

        elsif node.children[idx].is_a? Parser::AST::Node
          d[propname] = convert(node.children[idx])

        else
          d[propname] = node.children[idx]
        end
      end

      if cdr_index != nil and cdr_key != nil and node.respond_to?("children") and \
        node.children.length > cdr_index

        d[cdr_key] = node.children[cdr_index..-1].map{ |x| convert(x) }.compact
      end

      return add_position(node, d)
    end

    def add_position(node, hash)

      def add_from_subelem(node, hash, key)
        hash["start_line"] = node.loc.send(key).begin.line
        hash["end_line"] = node.loc.send(key).end.line
        hash["start_col"] = node.loc.send(key).begin.column
        hash["end_col"] = node.loc.send(key).end.column
      end

      case hash["type"]

      when "defined?", "module", "class", "sclass", "def", "defs",
        "undef", "alias", "super", "zsuper", "yield", "if", "when",
        "case", "while", "while-post", "for", "break", "next", "redo",
        "return", "rescue", "ensure", "retry", "preexe", "postexe"
        subelem = "keyword"

      when "optarg", "restarg", "blockarg", "kwarg", "kwoptarg",
           "kwrestarg"
        subelem = "name"

      when "not"
        subelem = "operator"

      else
        subelem = "expression"
      end

      add_from_subelem(node, hash, subelem)

      return hash
    end

  end
end

if __FILE__ == $0
  require 'pp'
  content = <<CODE
  123
CODE
  node = Parser::CurrentRuby.parse(content)
  pp(node)
  n = NodeConverter::Converter.new(node)
  pp(n.tohash)
end
