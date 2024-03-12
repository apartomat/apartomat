import React from "react"

import { Box, Layer, Text } from "grommet"
import { StatusGood } from "grommet-icons"

export function Notification({ message, ...layerProps }: { message: string }) {
    return (
        <Layer
            position="top"
            modal={false}
            responsive={false}
            margin={{ vertical: "small", horizontal: "small" }}
            {...layerProps}
        >
            <Box
                align="center"
                direction="row"
                gap="xsmall"
                justify="between"
                elevation="small"
                background="status-ok"
                round="medium"
                pad={{ vertical: "xsmall", horizontal: "small" }}
            >
                <StatusGood />
                <Text>{message}</Text>
            </Box>
        </Layer>
    )
}

export default Notification
