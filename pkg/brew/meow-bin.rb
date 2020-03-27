class MeowBin < Formula
  desc "Day previewer written in Golang"
  homepage "https://github.com/epylinkn/meow"
  url 'https://github.com/epylinkn/meow/releases/download/v0.1.0/meow-bin.tar.gz'
  sha256 '3dbbb512f2d9a82a706b6be471e13c8922a53c3d29d0a7627225667f1629c8ed'
  version '0.1.0'

  depends_on "ical_buddy"

  def install
    bin.install "meow"
  end
end
