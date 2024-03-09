import React from "react"

import { Spinner, SpinnerExtendedProps } from "grommet"

export default function Loading(props: SpinnerExtendedProps) {
    return (
        <Spinner
            border={[
                { side: "all", color: "background-contrast", size: "medium" },
                { side: "right", color: "brand", size: "medium" },
                { side: "top", color: "brand", size: "medium" },
                { side: "left", color: "brand", size: "medium" },
            ]}
            {...props}
        />
    )
}
