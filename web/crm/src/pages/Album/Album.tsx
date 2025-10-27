import React, { useEffect, useRef, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"

import { Box, BoxExtendedProps, Button, Grid, Heading, Header, Text, Drop, Main } from "grommet"
import { Add, Close } from "grommet-icons"

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
import { AddSplitCover } from "./AddSplitCover/AddSplitCover"
import { Page } from "pages/Album/Page"

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

    const [showAddSplitCover, setShowAddSplitCover] = useState(false)

    const [scale, setScale] = useState(1.0)

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
                    setScale(0.7)
                    break
            }
        }
    }, [data])

    const navigate = useNavigate()

    const [inView, setInView] = useState(0)

    const [hoverPage, setHoverPage] = useState<number | undefined>()

    const handlePageMouseEnter = (n: number) => {
        setHoverPage(n)
    }

    const handlePageMouseLeave = (n: number) => {
        setHoverPage(undefined)
    }

    return (
        <Main overflow="scroll" style={{ position: "fixed", inset: 0 }} background="background-contrast">
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
                        <PageSize albumId={id} size={settings.pageSize} onAlbumPageSizeChanged={() => refetch()} />
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
                onClickAddSplitCover={() => setShowAddSplitCover(true)}
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
                <Box align="center" pad={{ top: "90px", bottom: "68px" }}>
                    <Grid width="100%">
                        {pages.map((p, pageNumber) => {
                            return (
                                <Page
                                    key={pageNumber}
                                    pageNumber={pageNumber}
                                    albumId={id}
                                    page={p}
                                    settings={settings}
                                    scale={scale}
                                    onPageDeleted={() => {
                                        refetch()
                                    }}
                                />
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

            {showAddSplitCover && (
                <AddSplitCover
                    albumId={id}
                    onClickClose={() => setShowAddSplitCover(false)}
                    onSplitCoverAdded={async () => {
                        setShowAddSplitCover(false)
                        await refetch()
                    }}
                />
            )}
        </Main>
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
    onClickAddSplitCover,
    ...boxProps
}: {
    onClickAddVisualizations?: () => void
    onClickUploadCover?: () => void
    onClickAddSplitCover?: () => void
} & BoxExtendedProps) {
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
                            label="Загрузить обложку"
                            onClick={() => {
                                setOpen(false)
                                onClickUploadCover && onClickUploadCover()
                            }}
                        />
                        <Button
                            primary
                            label="Обложку"
                            color="accent-2"
                            onClick={() => {
                                setOpen(false)
                                onClickAddSplitCover && onClickAddSplitCover()
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
