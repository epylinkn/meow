class MeowBin < Formula
  desc "Day previewer written in Golang"
  homepage "https://github.com/epylinkn/meow"
  url ''
  sha256 ''
  version '0.1.0'

  depends_on "ical_buddy"

  def install
    bin.install "meow"
  end
end
