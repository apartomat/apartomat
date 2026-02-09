import { Box, Button, Text } from "grommet"

import { Key, useState } from "react"
import {
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    PageOrientation as PageOrientationEnum,
    PageSize as PageSizeEnum,
} from "api/graphql"
import { SplitCover } from "features/album/cover/SplitCover/SplitCover"
import { Edit } from "grommet-icons"

import { DeletePage, InnerHtmlPage } from "./ui"

export function Page({
    pageNumber,
    albumId,
    page,
    settings,
    scale,
    onPageDeleted,
}: {
    key: Key
    pageNumber: number
    albumId: string
    page: AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment
    settings: AlbumScreenSettingsFragment
    scale: number
    onPageDeleted: () => void
}) {
    const [hovered, setHovered] = useState(false)

    return (
        <Box
            direction="row"
            justify="center"
            onMouseOver={() => {
                setHovered(true)
            }}
            onMouseOut={() => {
                setHovered(false)
            }}
        >
            <Box direction="column" justify="center" gap="small">
                <Box
                    pad="xsmall"
                    style={{ visibility: hovered ? "visible" : "hidden" }}
                    background="background-contrast"
                    round="small"
                    direction="column"
                    gap="small"
                    margin={{ right: "xsmall" }}
                    alignSelf="end"
                >
                    <DeletePage key={page.id} albumId={albumId} pageNumber={pageNumber} onPageDeleted={onPageDeleted} />
                </Box>
            </Box>
            <Box
                background="background-front"
                margin={{ vertical: "xsmall" }}
                width={orientationWidth(settings.pageSize, settings.pageOrientation, scale)}
                height={orientationHeight(settings.pageSize, settings.pageOrientation, scale)}
                overflow="hidden"
            >
                {(() => {
                    switch (page.__typename) {
                        case "AlbumPageCover":
                            switch (page.cover.__typename) {
                                case "SplitCover":
                                    return (
                                        <SplitCover
                                            scale={scale}
                                            title={page.cover.title}
                                            subtitle={page.cover.subtitle}
                                            qrcodeSrc={page.cover.qrCodeSrc}
                                            city={page.cover.city}
                                            year={page.cover.year}
                                            image={page.cover.image}
                                            logoText="PUHOVA"
                                        />
                                    )
                                case "CoverUploaded":
                                    return (
                                        <InnerHtmlPage
                                            scale={scale}
                                            html={page.svg.__typename === "Svg" ? page.svg.svg : ""}
                                        />
                                    )
                                default:
                                    return <></>
                            }
                        case "AlbumPageVisualization":
                            switch (page.visualization.__typename) {
                                case "Visualization":
                                    return (
                                        <InnerHtmlPage
                                            scale={scale}
                                            html={page.svg.__typename === "Svg" ? page.svg.svg : ""}
                                        />
                                    )
                                default:
                                    return <></>
                            }
                        default:
                            return <></>
                    }
                })()}
            </Box>
            {false && (
                <Box direction="column" justify="center" gap="small">
                    <Box
                      pad="xsmall"
                      style={{ visibility: hovered ? "visible" : "hidden" }}
                      background="background-contrast"
                      round="small"
                      direction="column"
                      gap="small"
                      margin={{ left: "xsmall" }}
                      alignSelf="end"
                    >
                        <Button plain>
                            <Edit />
                        </Button>
                    </Box>
                </Box>
              )
            }
        </Box>
    )
}

function orientationWidth(size: PageSizeEnum, orientation: PageOrientationEnum, scale: number = 1.0): string {
    const width = {
        [PageSizeEnum.A3]: {
            [PageOrientationEnum.Portrait]: 297,
            [PageOrientationEnum.Landscape]: 420,
        },
        [PageSizeEnum.A4]: {
            [PageOrientationEnum.Portrait]: 210,
            [PageOrientationEnum.Landscape]: 297,
        },
    }

    return `${width[size][orientation] * scale}mm`
}

function orientationHeight(
    size: PageSizeEnum,
    orientation: PageOrientationEnum = PageOrientationEnum.Landscape,
    scale: number = 1.0
): string {
    const height = {
        [PageSizeEnum.A3]: {
            [PageOrientationEnum.Portrait]: 420,
            [PageOrientationEnum.Landscape]: 297,
        },
        [PageSizeEnum.A4]: {
            [PageOrientationEnum.Portrait]: 297,
            [PageOrientationEnum.Landscape]: 210,
        },
    }

    return `${height[size][orientation] * scale}mm`
}
