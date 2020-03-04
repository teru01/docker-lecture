ReVIEW::Compiler.definline :seqsplit
module ReVIEW
  class LATEXBuilder
    def inline_seqsplit(str)
      escaped_str = escape_latex(str)
      "\\seqsplit{#{escaped_str}}"
    end
  end
  class HTMLBuilder
    def inline_seqsplit(str)
      str
    end
  end
end
