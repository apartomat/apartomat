import { Box } from "grommet"
import { DeletePage } from "pages/Album/Page/ui/DeletePage"
import React, { Key, useState } from "react"
import {
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    PageOrientation as PageOrientationEnum,
    PageSize as PageSizeEnum,
} from "api/graphql"
import { SplitCover } from "features/album/cover/SplitCover/SplitCover"

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
            <Box direction="column" justify="center">
                <Box
                    pad="xsmall"
                    style={{ visibility: hovered ? "visible" : "hidden" }}
                    background="background-contrast"
                    round="small"
                    direction="column"
                    gap="small"
                    margin={{ right: "xsmall" }}
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
                                            imgSrc="http://localhost:9000/apartomat/p/K0UkDwTV8JUU792OKKojA/i2TtzpZ8F2D6Ax3hvpy43.jpg"
                                            logoText="PUHOVA"
                                        />
                                    )
                                case "CoverUploaded":
                                    return (
                                        <Box
                                            style={{
                                                transform: `scale(${scale})`,
                                                transformOrigin: "left top",
                                            }}
                                        >
                                            {page.svg.__typename === "Svg" && (
                                                <div dangerouslySetInnerHTML={svg(page.svg.svg)} />
                                            )}
                                        </Box>
                                    )
                                default:
                                    return <></>
                            }
                        case "AlbumPageVisualization":
                            switch (page.visualization.__typename) {
                                case "Visualization":
                                    return (
                                        <Box
                                            style={{
                                                transform: `scale(${scale})`,
                                                transformOrigin: "left top",
                                            }}
                                        >
                                            {page.svg.__typename === "Svg" && (
                                                <div dangerouslySetInnerHTML={svg(page.svg.svg)} />
                                            )}
                                        </Box>
                                    )
                                default:
                                    return <></>
                            }
                        default:
                            return <></>
                    }
                })()}
            </Box>
            <Box width="xxsmall"></Box>
        </Box>
    )
}

function svg(html: string) {
    return { __html: html }
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
