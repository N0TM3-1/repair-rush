components {
  id: "player"
  component: "/scripts/player.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"placeholder_player\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlasses/player.atlas\"\n"
  "}\n"
  ""
  scale {
    x: 5.0
    y: 5.0
  }
}
