import React from "react"

import { Box, BoxExtendedProps, Image, Text } from "grommet"

import { ProjectScreenVisualizationsFragment as ProjectScreenVisualizations } from "api/types"

export default function Visualizations({
    visualizations,
    ...boxProps
}: {
    visualizations: ProjectScreenVisualizations
} & BoxExtendedProps) {
    switch (visualizations.list.__typename) {
        case "ProjectVisualizationsList":
            if (visualizations.list.items.length === 0) {
                return null
            }

            return (
                <Box {...boxProps}>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}}>
                        {visualizations.list.items.map(vis =>
                            <Box
                                key={vis.id}
                                height="small"
                                width="small"
                                margin={{bottom: "small"}}
                                flex={{"shrink":0}}
                                background="light-2"
                            >
                                <Image
                                    fit="cover"
                                    src={`${vis.file.url}?w=192`}
                                    srcSet={`${vis.file.url}?w=192 192w, ${vis.file.url}?w=384 384w`}
                                />
                            </Box>
                        )}
                        {visualizations.total.__typename === "ProjectVisualizationsTotal" && visualizations.total.total > visualizations.list.items.length
                            ? <Box key={0} height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}} align="center" justify="center">
                                <Text>ещё {visualizations.total.total - visualizations.list.items.length}</Text>
                            </Box>
                            : null
                        }
                    </Box>
                </Box>
            )
        default:
            return null
    }
}