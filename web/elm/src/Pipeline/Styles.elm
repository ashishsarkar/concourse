module Pipeline.Styles exposing
    ( cliIcon
    , groupItem
    , groupsBar
    , pauseToggle
    )

import Colors
import Concourse.Cli as Cli
import Html
import Html.Attributes exposing (style)


groupsBar : List (Html.Attribute msg)
groupsBar =
    [ style "background-color" Colors.groupsBarBackground
    , style "color" Colors.dashboardText
    , style "display" "flex"
    , style "flex-flow" "row wrap"
    , style "padding" "5px"
    ]


groupItem : { selected : Bool, hovered : Bool } -> List (Html.Attribute msg)
groupItem { selected, hovered } =
    [ style "font-size" "14px"
    , style "background" Colors.groupBackground
    , style "margin" "5px"
    , style "padding" "10px"
    ]
        ++ (if selected then
                [ style "opacity" "1"
                , style "border" <| "1px solid " ++ Colors.groupBorderSelected
                ]

            else if hovered then
                [ style "opacity" "0.6"
                , style "border" <| "1px solid " ++ Colors.groupBorderHovered
                ]

            else
                [ style "opacity" "0.6"
                , style "border" <| "1px solid " ++ Colors.groupBorderUnselected
                ]
           )


pauseToggle : Bool -> List (Html.Attribute msg)
pauseToggle isPaused =
    [ style "border-left" <|
        if isPaused then
            "1px solid rgba(255, 255, 255, 0.5)"

        else
            "1px solid #3d3c3c"
    ]


cliIcon : Cli.Cli -> List (Html.Attribute msg)
cliIcon cli =
    [ style "width" "12px"
    , style "height" "12px"
    , style "background-image" <| Cli.iconUrl cli
    , style "background-repeat" "no-repeat"
    , style "background-position" "50% 50%"
    , style "background-size" "contain"
    , style "display" "inline-block"
    ]
