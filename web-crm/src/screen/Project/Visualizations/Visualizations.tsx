import React, { useRef, useEffect, useState } from "react"

import { ProjectScreenVisualizationsFragment as ProjectScreenVisualizations } from "api/types"

import { Box, BoxExtendedProps, Grid, Image, Text } from "grommet"
import { Next, Previous } from "grommet-icons"

export default function Visualizations({
    visualizations,
    ...boxProps
}: {
    visualizations: ProjectScreenVisualizations
} & BoxExtendedProps) {

    const gridRef = useRef<HTMLDivElement>(null)

    const [ left, setLeft ] = useState<boolean>(true)

    const [ right, setRight ] = useState<boolean>(false)

    useEffect(() => {
        const grid = gridRef.current

        const handle = () => {
            if (grid) {
                setLeft(grid.scrollLeft === 0)
                setRight(grid.scrollLeft + grid.offsetWidth === grid.scrollWidth)
            }
        }

        grid?.addEventListener("scroll", handle)

        return () => {
            grid?.removeEventListener("scroll", handle)
        }
    })

    switch (visualizations.list.__typename) {
        case "ProjectVisualizationsList":
            if (visualizations.list.items.length === 0) {
                return null
            }

            return (
                <Box {...boxProps} direction="row">
                    {!left &&
                        <Box align="end" justify="center" height="small" width="xxsmall" style={{ position: "absolute", left: 0 }}>
                            <Previous color="light-4"/>
                        </Box>
                    }

                    <Box overflow="auto">
                        <Grid
                            columns="small"
                            style={{gridAutoFlow: "column", overflowX: "scroll"}}
                            gap="xsmall"
                            ref={gridRef}
                        >
                            {visualizations.list.items.map((item) => (
                                <Box
                                    key={item.id}
                                    height="small"
                                    width="small"
                                    flex={{"shrink":0}}
                                    background="light-2"
                                >
                                    <Image
                                        fit="cover"
                                        src={`${item.file.url}?w=192`}
                                        srcSet={`${item.file.url}?w=192 192w, ${item.file.url}?w=384 384w`}
                                    />
                                </Box>
                            ))}
                        </Grid>

                        {visualizations.total.__typename === "ProjectVisualizationsTotal" && visualizations.total.total > visualizations.list.items.length
                            ? <Box key={0} height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}} align="center" justify="center">
                                <Text>ещё {visualizations.total.total - visualizations.list.items.length}</Text>
                            </Box>
                            : null
                        }
                    </Box>
                    {!right &&
                        <Box align="start" justify="center" height="small" width="xxsmall" style={{ position: "absolute", right: 0 }}>
                            <Next color="light-3"/>
                        </Box>
                    }
                </Box>
            )
        default:
            return null
    }
}