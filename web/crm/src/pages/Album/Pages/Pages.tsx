import { AlbumScreenAlbumPageCoverFragment, AlbumScreenAlbumPageVisualizationFragment } from "api/graphql"
import { Box, BoxExtendedProps, Grid, Image } from "grommet"
import Paper from "pages/Album/Pages/Paper/Paper"

export function Pages({
    pages,
    current,
    onClickPage,
    ...props
}: {
    pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]
    current: number
    onClickPage: (n: number) => void
} & BoxExtendedProps) {
    return (
        <Box overflow="auto" pad="xsmall" {...props}>
            <Grid columns="xsmall" style={{ gridAutoFlow: "column", overflowX: "scroll" }} gap="xsmall" pad="xsmall">
                {pages.map((p, key) => {
                    return (
                        <Box
                            key={key}
                            height="xsmall"
                            width="xsmall"
                            flex={{ shrink: 0 }}
                            style={{ boxShadow: current === key ? "0 0 0px 2px #7D4CDB" : "none" }}
                            align="center"
                        >
                            <Paper scale={0.1}>
                                {(() => {
                                    switch (p.__typename) {
                                        case "AlbumPageVisualization":
                                            switch (p.visualization.__typename) {
                                                case "Visualization":
                                                    return (
                                                        <Image
                                                            fit="cover"
                                                            src={p.visualization.file.url}
                                                            onClick={() => {
                                                                onClickPage && onClickPage(key)
                                                            }}
                                                        />
                                                    )
                                                default:
                                                    return <></>
                                            }
                                        default:
                                            return <></>
                                    }
                                })()}
                            </Paper>
                        </Box>
                    )
                })}
            </Grid>
        </Box>
    )
}
