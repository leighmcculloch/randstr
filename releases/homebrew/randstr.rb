class Randstr < Formula
  homepage "https://github.com/leighmcculloch/randstr"
  url "https://raw.githubusercontent.com/leighmcculloch/randstr/binaries/darwin/amd64/randstr"
  version "0.0.0"
  sha256 ""

  def install
    bin.install "randstr"
  end

  test do
    system "#{bin}/randstr", "-version"
  end
end
