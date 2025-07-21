components {
  id: "stove"
  component: "/scripts/machines/stove.script"
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"stove\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/atlasses/machines.atlas\"\n"
  "}\n"
  ""
}
