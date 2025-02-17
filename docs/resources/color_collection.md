---
page_title: "looker_color_collection Resource - terraform-provider-looker"
subcategory: ""
description: |-
  
---
# looker_color_collection (Resource)

## Example Usage
```terraform
resource "looker_color_collection" "collection" {
  label      = "My new collection"
  categoricalpalettes {
    label = "cat"
    colors = ["#1A73E8",
        "#12B5CB",
        "#E52592"]
  }
  sequentialpalettes {
    label = "seq"
    stops {
      color = "#FFFFFF"
      offset = "1"
    }
    stops {
      color = "#1A73E8"
      offset = "100"
    }
  }
  divergingpalettes {
    label = "div"
    stops {
      color = "#FFFFFF"
      offset = "1"
    }
    stops {
      color = "#1A73E8"
      offset = "100"
    }
  }
}
```

## Example Output
```terraform
# looker_color_collection.collection:
resource "looker_color_collection" "collection" {
    id    = "my-new-collection"
    label = "My new collection"

    categoricalpalettes {
        colors = [
            "#12B5CB",
            "#1A73E8",
            "#E52592",
        ]
        id     = "my-new-collection-categorical-0"
        label  = "cat"
        type   = "Categorical"
    }

    divergingpalettes {
        id    = "my-new-collection-diverging-0"
        label = "div"
        type  = "Diverging"

        stops {
            color  = "#1A73E8"
            offset = 100
        }
        stops {
            color  = "#FFFFFF"
            offset = 1
        }
    }

    sequentialpalettes {
        id    = "my-new-collection-sequential-0"
        label = "seq"
        type  = "Sequential"

        stops {
            color  = "#1A73E8"
            offset = 100
        }
        stops {
            color  = "#FFFFFF"
            offset = 1
        }
    }
}
```

<!-- schema generated by tfplugindocs -->
## Schema

### Required

- `categoricalpalettes` (Block Set, Min: 1) Array of categorical palette definitions (see [below for nested schema](#nestedblock--categoricalpalettes))
- `divergingpalettes` (Block Set, Min: 1) Array of categorical palette definitions (see [below for nested schema](#nestedblock--divergingpalettes))
- `label` (String) Label of color collection
- `sequentialpalettes` (Block Set, Min: 1) Array of categorical palette definitions (see [below for nested schema](#nestedblock--sequentialpalettes))

### Read-Only

- `id` (String) ColorCollection id

<a id="nestedblock--categoricalpalettes"></a>
### Nested Schema for `categoricalpalettes`

Required:

- `colors` (Set of String)

Optional:

- `label` (String) Label of palette
- `type` (String) Type of palette

Read-Only:

- `id` (String) Unique identity string


<a id="nestedblock--divergingpalettes"></a>
### Nested Schema for `divergingpalettes`

Required:

- `stops` (Block Set, Min: 2) Array of ColorStops in the palette (see [below for nested schema](#nestedblock--divergingpalettes--stops))

Optional:

- `label` (String) Label for palette
- `type` (String) Type of palette

Read-Only:

- `id` (String) Unique identity string

<a id="nestedblock--divergingpalettes--stops"></a>
### Nested Schema for `divergingpalettes.stops`

Required:

- `color` (String) CSS color string
- `offset` (Number) Offset in continuous palette (0 to 100)



<a id="nestedblock--sequentialpalettes"></a>
### Nested Schema for `sequentialpalettes`

Required:

- `stops` (Block Set, Min: 2) Array of ColorStops in the palette (see [below for nested schema](#nestedblock--sequentialpalettes--stops))

Optional:

- `label` (String) Label of palette
- `type` (String) type of palette

Read-Only:

- `id` (String) Unique ID of palette

<a id="nestedblock--sequentialpalettes--stops"></a>
### Nested Schema for `sequentialpalettes.stops`

Required:

- `color` (String) CSS color string
- `offset` (Number) Offset in continuous palette (0 to 100)
