import { useRef, useEffect, useState } from "react"
import { useNavigate } from "react-router-dom"

import { ProjectScreenVisualizationsFragment as ProjectScreenVisualizations } from "api/graphql"

import { Box, BoxExtendedProps, Grid, Image, Button } from "grommet"
import { Next, Previous } from "grommet-icons"

export default function Visualizations({
    projectId,
    visualizations,
    ...boxProps
}: {
    projectId: string
    visualizations: ProjectScreenVisualizations
} & BoxExtendedProps) {
    const gridRef = useRef<HTMLDivElement>(null)

    const [left, setLeft] = useState<boolean>(true)

    const [right, setRight] = useState<boolean>(false)

    const navigate = useNavigate()

    useEffect(() => {
        const grid = gridRef.current

        const handleScroll = () => {
            if (grid) {
                setLeft(grid.scrollLeft === 0)
                setRight(grid.scrollLeft + grid.offsetWidth === grid.scrollWidth)
            }
        }

        grid?.addEventListener("scroll", handleScroll)

        return () => {
            grid?.removeEventListener("scroll", handleScroll)
        }
    })

    const handleClickMore = () => {
        navigate(`/vis/${projectId}`)
    }

    switch (visualizations.list.__typename) {
        case "ProjectVisualizationsList":
            if (visualizations.list.items.length === 0) {
                return null
            }

            return (
                <Box {...boxProps} direction="row">
                    {!left && (
                        <Box
                            align="end"
                            justify="center"
                            height="small"
                            width="xxsmall"
                            style={{ position: "absolute", left: 0 }}
                        >
                            <Previous color="light-4" />
                        </Box>
                    )}

                    <Box overflow="auto">
                        <Grid
                            columns="small"
                            style={{ gridAutoFlow: "column", overflowX: "scroll" }}
                            gap="xsmall"
                            ref={gridRef}
                        >
                            {visualizations.list.items.map((item) => (
                                <Box
                                    key={item.id}
                                    height="small"
                                    width="small"
                                    flex={{ shrink: 0 }}
                                    background="light-2"
                                >
                                    <Image
                                        fit="cover"
                                        src={`${item.file.url}?w=192`}
                                        srcSet={`${item.file.url}?w=192 192w, ${item.file.url}?w=384 384w`}
                                    />
                                </Box>
                            ))}

                            {needMore(visualizations) && (
                                <Box
                                    key={0}
                                    height="small"
                                    width="xsmall"
                                    flex={{ shrink: 0 }}
                                    align="center"
                                    justify="center"
                                >
                                    <Button
                                        label={`ещё ${more(visualizations)}`}
                                        size="small"
                                        primary
                                        color="light-2"
                                        // icon={<Next size="small"/>}
                                        reverse
                                        onClick={handleClickMore}
                                    />
                                </Box>
                            )}
                        </Grid>
                    </Box>
                    {!right && (
                        <Box
                            align="start"
                            justify="center"
                            height="small"
                            width="xxsmall"
                            style={{ position: "absolute", right: 0 }}
                        >
                            <Next color="light-3" />
                        </Box>
                    )}
                </Box>
            )
        default:
            return null
    }
}

function needMore(vis: ProjectScreenVisualizations): boolean {
    if (vis.list.__typename !== "ProjectVisualizationsList") {
        return false
    }

    if (vis.total.__typename !== "ProjectVisualizationsTotal") {
        return false
    }

    return vis.total.total > vis.list.items.length
}

function more(vis: ProjectScreenVisualizations): number {
    if (vis.list.__typename !== "ProjectVisualizationsList") {
        return 0
    }

    if (vis.total.__typename !== "ProjectVisualizationsTotal") {
        return 0
    }

    return vis.total.total - vis.list.items.length
}
