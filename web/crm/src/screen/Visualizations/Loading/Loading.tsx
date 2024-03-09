import React from "react"

import { Spinner, SpinnerExtendedProps } from "grommet"
import { ColorType } from "grommet/utils"

export const Loading = ({
    color = "brand",
    weight = "medium",
    ...props
}: {
    color?: ColorType
    weight?: "xsmall" | "small" | "medium" | "large" | "xlarge"
} & SpinnerExtendedProps) => {
    return (
        <Spinner
            border={[
                { side: "all", color: "background-contrast", size: weight },
                { side: "right", color, size: weight },
                { side: "top", color, size: weight },
                { side: "left", color, size: weight },
            ]}
            {...props}
        />
    )
}

export default Loading
