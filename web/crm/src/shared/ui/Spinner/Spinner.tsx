import { Spinner as GrommetSpinner, SpinnerExtendedProps } from "grommet"
import { ColorType } from "grommet/utils"

export const Spinner = ({
    color = "brand",
    weight = "medium",
    ...props
}: {
    color?: ColorType
    weight?: "xsmall" | "small" | "medium" | "large" | "xlarge"
} & SpinnerExtendedProps) => {
    return (
        <GrommetSpinner
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
