import React, { MouseEventHandler, useEffect, useRef, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"
import { InView } from "react-intersection-observer"

import { Box, BoxExtendedProps, Button, Grid, Heading, Header, Text, Drop, Main } from "grommet"
import {Add, Close} from "grommet-icons"

import useAlbum, {
    AlbumScreenVisualizationFragment,
    AlbumScreenProjectFragment,
    AlbumScreenAlbumPageCoverFragment,
    AlbumScreenAlbumPageVisualizationFragment,
    AlbumScreenSettingsFragment,
    PageSize as PageSizeEnum,
    PageOrientation as PageOrientationEnum,
    AlbumScreenHouseRoomFragment,
} from "./useAlbum"

import { PageSize, PageOrientation } from "./Settings/"
import AddVisualizations from "pages/Album/AddVisualizations/AddVisualizations"
import { GenerateFile as GenerateAlbumFile } from "pages/Album/GenerateFile"
import { UploadCover } from "pages/Album/UploadCover"
import { DeletePage } from "pages/Album/DeletePage/";
import {Confirm} from "widgets/confirm";

export function Album() {
    const { id } = useParams<"id">() as { id: string }

    const [project, setProject] = useState<AlbumScreenProjectFragment | undefined>()

    const [visualizations, setVisualizations] = useState<AlbumScreenVisualizationFragment[]>([])

    const [rooms, setRooms] = useState<AlbumScreenHouseRoomFragment[]>([])

    const [pages, setPages] = useState<
        (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]
    >([])

    const [settings, setSettings] = useState<AlbumScreenSettingsFragment | undefined>()

    const { data, loading, refetch } = useAlbum({ id })

    const [showAddVisualizations, setShowAddVisualizations] = useState(false)

    const [showUploadCover, setShowUploadCover] = useState(false)

    const [scale, setScale] = useState(0.75)

    useEffect(() => {
        if (data?.album?.__typename === "Album") {
            if (data?.album?.project?.__typename === "Project") {
                setProject(data.album.project)

                if (data.album.project.visualizations.list.__typename === "ProjectVisualizationsList") {
                    setVisualizations(data.album.project.visualizations.list.items)
                }

                if (
                    data.album.project.houses.__typename === "ProjectHouses" &&
                    data.album.project.houses.list.__typename === "ProjectHousesList"
                ) {
                    const list = data.album.project.houses.list.items[0].rooms.list

                    if (list.__typename === "HouseRoomsList") {
                        setRooms(list.items)
                    }
                }
            }

            if (data.album.pages.__typename === "AlbumPages") {
                setPages(data.album.pages.items)
            }

            setSettings(data.album.settings)

            switch (data.album.settings.pageSize) {
                case PageSizeEnum.A3:
                    setScale(0.5)
                    break
                case PageSizeEnum.A4:
                    setScale(0.707)
                    break
            }
        }
    }, [data])

    const navigate = useNavigate()

    const [inView, setInView] = useState(0)

    const [ hoverPage, setHoverPage ] = useState<number | undefined>()

    const handlePageMouseEnter = (n: number) => {
        setHoverPage(n)
    }

    const handlePageMouseLeave = (n: number) => {
        setHoverPage(undefined)
    }

    return (
        <Main overflow="scroll" style={{ position: "fixed", inset: 0 }}>
            <Header
                style={{
                    position: "fixed",
                    top: 0,
                    left: 0,
                    right: 0,
                    zIndex: 2,
                }}
                background={{
                    color: "white",
                    opacity: "strong",
                }}
                pad={{ top: "medium", bottom: "small", horizontal: "large" }}
                justify="between"
            >
                <Grid columns={{ count: 3, size: "auto" }} gap="small" width="100%">
                    <Box align="start" justify="center">
                        <Text size="xlarge" weight="bold">
                            {data?.album.__typename === "Album" && <>{data?.album.name}</>}
                        </Text>
                    </Box>
                    <Box align="center" justify="center">
                        {pages.length > 0 && (
                            <Text weight="bold" color="text-xweak">
                                {inView + 1} / {pages.length}
                            </Text>
                        )}
                    </Box>
                    <Box align="end" justify="center">
                        <Button
                            icon={<Close />}
                            onClick={() => {
                                project && navigate(`/p/${project.id}`)
                            }}
                        />
                    </Box>
                </Grid>
            </Header>

            <Box
                style={{
                    position: "fixed",
                    top: "84px",
                    right: "60px",
                    zIndex: 1,
                }}
                background="background-back"
                pad="medium"
                round="xsmall"
                margin={{ top: "large" }}
            >
                {settings && (
                    <Box gap="small" align="end">
                        <Heading level={5} margin={{ top: "none" }}>
                            Настройки для печати
                        </Heading>
                        <PageSize
                            albumId={id}
                            size={settings.pageSize}
                            onAlbumPageSizeChanged={() => refetch()}
                        />
                        <PageOrientation
                            albumId={id}
                            orientation={settings.pageOrientation}
                            onAlbumPageOrientationChanged={() => refetch()}
                        />
                    </Box>
                )}
                {data?.album.__typename === "Album" && (
                    <Box gap="small" align="end" margin={{ top: "medium" }}>
                        <GenerateAlbumFile album={data.album} onAlbumFileGenerated={() => refetch()} />
                    </Box>
                )}
            </Box>

            <AddVariants
                onClickAddVisualizations={() => setShowAddVisualizations(true)}
                onClickUploadCover={() => setShowUploadCover(true)}
                style={{
                    position: "fixed",
                    bottom: 0,
                    left: 0,
                    right: 0,
                    zIndex: 1,
                }}
                direction="row"
                justify="center"
                pad="small"
            />

            {pages.length > 0 && settings && (
                <Box align="center" pad={{ top: "84px", bottom: "68px" }}>
                    <Grid width="100%">
                        {pages.map((p, key) => {
                            return (
                                <Box
                                    key={key}
                                    direction="row"
                                    justify="center"
                                    onMouseOver={() => {
                                        handlePageMouseEnter(key)

                                    }}
                                    onMouseOut={() => {
                                        handlePageMouseLeave(key)
                                    }}
                                >
                                    <Box direction="column" justify="center">
                                        <Box
                                            pad="xsmall"
                                            style={{ visibility: hoverPage === key ? "visible": "hidden" }}
                                            background="background-contrast"
                                            round="small"
                                            direction="column"
                                            gap="small"
                                            margin={{right: "xsmall"}}
                                        >
                                            <DeletePage
                                                key={p.id}
                                                albumId={id}
                                                pageNumber={key}
                                                onPageDeleted={() => {
                                                    refetch()
                                                }}
                                            />
                                        </Box>
                                    </Box>
                                    <Box
                                        background="background-contrast"
                                        margin={{ vertical: "xxsmall" }}
                                        width={orientationWidth(settings.pageSize, settings.pageOrientation, scale)}
                                        height={orientationHeight(settings.pageSize, settings.pageOrientation, scale)}

                                    >
                                        {(() => {
                                            switch (p.__typename) {
                                                case "AlbumPageVisualization":
                                                    switch (p.visualization.__typename) {
                                                        case "Visualization":
                                                            return (
                                                                <Box
                                                                    style={{
                                                                        transform: `scale(${scale})`,
                                                                        transformOrigin: "left top",
                                                                    }}
                                                                >
                                                                    {p.svg.__typename === "Svg" && (
                                                                        <div dangerouslySetInnerHTML={svg(p.svg.svg)} />
                                                                    )}
                                                                </Box>
                                                            )
                                                        default:
                                                            return <></>
                                                    }
                                                case "AlbumPageCover":
                                                    switch (p.cover.__typename) {
                                                        case "CoverUploaded":
                                                            return (
                                                                <Box
                                                                    style={{
                                                                        transform: `scale(${scale})`,
                                                                        transformOrigin: "left top",
                                                                    }}
                                                                >
                                                                    {p.svg.__typename === "Svg" && (
                                                                        <div dangerouslySetInnerHTML={svg(p.svg.svg)} />
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
                        })}
                    </Grid>
                </Box>
            )}

            {showAddVisualizations && (
                <AddVisualizations
                    albumId={id}
                    visualizations={visualizations}
                    rooms={rooms}
                    alreadyAdded={ids(pages)}
                    onVisualizationsAdded={() => {
                        setShowAddVisualizations(false)
                        refetch()
                    }}
                    onEsc={() => setShowAddVisualizations(false)}
                    onClickOutside={() => setShowAddVisualizations(false)}
                    onClickClose={() => setShowAddVisualizations(false)}
                />
            )}

            {showUploadCover && (
                <UploadCover
                    albumId={id}
                    onClickClose={() => setShowUploadCover(false)}
                    onAlbumCoverUploaded={async () => {
                        setShowUploadCover(false)
                        await refetch()
                    }}
                />
            )}
        </Main>
    )
}

function orientationWidth(
    size: PageSizeEnum,
    orientation: PageOrientationEnum,
    scale: number = 1.0
): string {
    const width = {
        [PageSizeEnum.A3]: {
            [PageOrientationEnum.Portrait]: 297,
            [PageOrientationEnum.Landscape]: 420,

        },
        [PageSizeEnum.A4]: {
            [PageOrientationEnum.Portrait]: 210,
            [PageOrientationEnum.Landscape]: 297,
        }
    }

    return `${width[size][orientation] * scale}mm`;
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
        }
    }

    return `${height[size][orientation] * scale}mm`;
}

/* eslint-disable  @typescript-eslint/no-explicit-any */
function ids(pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]): string[] {
    const vis = pages.filter(
        (p) => p.__typename === "AlbumPageVisualization" && p.visualization.__typename === "Visualization"
    ) as { visualization: { __typename?: "Visualization"; id: string; file: { __typename?: "File"; url: any } } }[]

    return vis.map((v) => v.visualization.id)
}

function AddVariants({
    onClickAddVisualizations,
    onClickUploadCover,
    ...boxProps
}: {
    onClickAddVisualizations?: () => void
    onClickUploadCover?: () => void
} & { boxProps: BoxExtendedProps }) {
    const [open, setOpen] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box {...boxProps}>
            <Box ref={targetRef} border={{ color: "background-front", size: "medium" }} round="large">
                <Button label="Добавить..." icon={<Add />} primary onClick={() => setOpen(true)} />
            </Box>

            {open && targetRef.current && (
                <Drop
                    elevation="none"
                    target={targetRef.current}
                    onClickOutside={() => setOpen(false)}
                    onEsc={() => setOpen(false)}
                    align={{ bottom: "bottom" }}
                    round="large"
                >
                    <Box gap="small" border={{ color: "background-front", size: "medium" }} direction="row">
                        <Button
                            primary
                            label="Обложку"
                            onClick={() => {
                                setOpen(false)
                                onClickUploadCover && onClickUploadCover()
                            }}
                        />
                        <Button
                            primary
                            label="Визуализации"
                            color="accent-3"
                            onClick={() => {
                                setOpen(false)
                                onClickAddVisualizations && onClickAddVisualizations()
                            }}
                        />
                        <Button primary label="Ссылку на сайт" color="status-ok" />
                    </Box>
                </Drop>
            )}
        </Box>
    )
}

function svg(html: string) {
    return { __html: html }
}
