{ pkgs }:
with pkgs;
buildGoModule rec {
  pname = "legato";
  version = "main";

  src = ../.;

  buildInputs = [ vips ];

  # This needs to be updated every time go dependencies change
  vendorSha256 = "sha256-KfcyCWE/VKdrKbtI6igvoDHtCPsCmOKCH+7F0qZzI2Q=";

  subPackages = [ "." ];

  meta = with pkgs.lib; {
    description = "Reference server implementation for the Harmony protocol.";
    homepage = "https://github.com/harmony-development/legato";
    license = licenses.agpl3;
    maintainers = [ maintainers.yusdacra ];
    platforms = platforms.linux;
  };
}
