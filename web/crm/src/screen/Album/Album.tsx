import React, { useEffect, useRef, useState } from "react"
import { useParams, useNavigate } from "react-router-dom"

import { Box, BoxExtendedProps, Button, Carousel, Drop, Grid, Heading, Image, Main } from "grommet"
import { Close, Previous, Next, Add, Sort } from "grommet-icons"
import useAlbum, { AlbumScreenVisualizationFragment, AlbumScreenProjectFragment, AlbumScreenAlbumPageCoverFragment, AlbumScreenAlbumPageVisualizationFragment, AlbumScreenSettingsFragment } from "./useAlbum"
import { useAddVisualizationsToAlbum } from "./useAddVisualizationsToAlbum"

import { PageSize, PageOrientation } from "./Settings/"

export function Album() {
    const { id } = useParams<"id">() as { id: string }

    const [ project, setProject ] = useState<AlbumScreenProjectFragment | undefined>()

    const [ visualizations, setVisualizations ] = useState<AlbumScreenVisualizationFragment[]>([])

    const [ pages, setPages ] = useState<(AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[]>([])

    const [ addVisualizations, { loading: addVisualizationsLoading } ] = useAddVisualizationsToAlbum(id)

    const { data } = useAlbum({ id })

    const [ settings, setSettings ] = useState<AlbumScreenSettingsFragment | undefined>()

    useEffect(() => {
        if (data?.album?.__typename === "Album") {
            if (data?.album?.project?.__typename === "Project") {
                setProject(data.album.project)

                if (data.album.project.visualizations.list.__typename === "ProjectVisualizationsList") {
                    setVisualizations(data.album.project.visualizations.list.items)
                }
            }    
            
            if (data.album.pages.__typename === "AlbumPages") {
                setPages(data.album.pages.items)
            }

            setSettings(data.album.settings)
        }
        
    }, [ data ])

    const [ currentPage, setCurrentPage ] = useState<number>(0)

    const [ redirectTo, setRedirectTo ] = useState<string | undefined>(undefined)

    const navigate = useNavigate()

    if (redirectTo) {
        navigate(redirectTo)
    }

    return (
        <Main pad={{vertical: "medium", horizontal: "large"}} direction="row" justify="center">
            <Box style={{ position: "fixed", right: 0 }} gap="small" margin={{ horizontal: "large" }}>
                <Box align="end">
                    <Button
                        icon={<Close/>}
                        onClick={() => {
                            project && setRedirectTo(`/p/${project.id}`)
                        } }
                    />
                </Box>
                {settings &&
                    <Box gap="small" align="end" margin={{ top: "large" }}>
                        <Heading level={5}>Настройки для печати</Heading>
                        <PageSize albumId={id} size={settings.pageSize} />
                        <PageOrientation albumId={id} orientation={settings.pageOrientation} />
                    </Box>
                }

            </Box>

            <Box style={{ position: "fixed", left: 0 }} gap="small" margin={{ horizontal: "large" }}>
                {data?.album.__typename === "Album" && <Heading level={3} margin={{top: ""}}>{data?.album.name}</Heading>}
            </Box>

            <Box direction="column" justify="center" height="large">
                    <Carousel
                        controls={false}
                        activeChild={currentPage}
                        onChild={(n) => {
                            setCurrentPage(n)
                        }}
                        fill
                    >
                        {pages.map((p, key) => {
                            return (
                                <Paper
                                    key={key}
                                        size="A4"
                                        scale={0.6}
                                        pad="medium"
                                        background="background-contrast"
                                        round="xsmall"
                                        justify="center"
                                        margin="xsmall"
                                    >
                                    {(() => {
                                            switch (p.__typename) {
                                                case "AlbumPageVisualization":
                                                    switch (p.visualization.__typename) {
                                                        case "Visualization":
                                                            return (
                                                                <Image
                                                                    fit="cover"
                                                                    src={p.visualization.file.url}
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
                            )
                        })}
                    </Carousel>
                {/* <Paper size="A4" scale={0.6} pad="medium" background="background-contrast">
                    {currentPage &&
                        <Image src={currentPage} fit="contain"/>
                    }
                </Paper> */}
            </Box>

            <Box
                direction="row"
                style={{ position: "absolute", bottom: 0 }}
                align="center"
            >
                <Box align="center" margin={{right: "large"}}>
                    <Box round="full" overflow="hidden" background="light-2">
                        <Button icon={<Sort/>} hoverIndicator/>
                    </Box>
                </Box>

                {/* {pages.length > 0 && <Box align="end"><Previous color="light-3"/></Box>} */}
                
                {pages.length > 0 && <Pages
                    pages={pages}
                    current={currentPage}
                    onClickPage={(n) => setCurrentPage(n)}
                    width={{max: "large"}}
                />}

                {/* {pages.length > 0 && <Box><Next color="light-3"/></Box>} */}

                <AddVisualizations
                    margin={{ left: "large" }}
                    visualizations={visualizations}
                    onClickAdd={(ids: string[]) => {
                        addVisualizations(ids)
                    }}
                />
            </Box>
        </Main>
    )
}

export default Album

function Paper({
    children,
    size = "A4",
    scale = 1.0,
    ...boxProps
}: {
    children: JSX.Element | never[] | undefined | string,
    size?: "A4" | "A5",
    scale?: 0.05 | 0.1 | 0.25 | 0.3 | 0.4 | 0.5 | 0.6 | 0.7 | 0.75 | 1
} & BoxExtendedProps) {
    return (
        <Box {...boxProps} width={`calc(${scale} * 210mm)`} height={`calc(${scale} * 297mm)`}>
            {children}
        </Box>
    )
}

function AddVisualizations({
    visualizations,
    onClickAdd,
    ...boxProps
}: {
    visualizations: AlbumScreenVisualizationFragment[],
    onClickAdd?: (id: string[]) => void 
} & BoxExtendedProps) {

    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    const [ selected, setSelected ] = useState<string[]>([])

    const select = (id: string) => {
        if (selected.includes(id)) {
            setSelected(selected.filter(s => s !== id))
        } else {
            setSelected([...selected, id])
        }
    }

    return (
        <Box {...boxProps} ref={targetRef}>
            <Box round="full" overflow="hidden" background="brand">
                <Button
                    icon={<Add />}
                    hoverIndicator
                    onClick={() => {
                        setShow(!show)
                        setSelected([])
                    }}/>
            </Box>
            {show && targetRef.current &&
                <Drop
                    target={targetRef.current}
                    align={{left: "right", top: "top"}}
                    elevation="small"
                    onClickOutside={() => {
                        setShow(false)
                        setSelected([])
                    }}
                    round="xsmall"
                >
                    {visualizations &&
                        <Box width="large" pad="small">
                            <Grid
                                columns="xxsmall"
                                rows="xxsmall"
                                gap="xsmall"
                            >
                                {visualizations.map((vis, key) => {
                                    return (
                                        <Box
                                            key={key}
                                            focusIndicator={false}
                                            round="xsmall"
                                            pad="xxsmall"
                                            style={{boxShadow: selected.includes(vis.id) ? "0 0 0px 2px #7D4CDB": "none" }}
                                            onClick={() => select(vis.id)}
                                        >
                                            <Image
                                                src={`${vis.file.url}?w=48`}
                                                fit="contain"
                                            ></Image>
                                        </Box>
                                    )
                                })}
                            </Grid>
                        </Box>
                    }
                    <Box direction="row" justify="between">
                        <Button
                            label="Добавить"
                            disabled={selected.length === 0}
                            onClick={() => {
                                onClickAdd && onClickAdd(selected)
                                setShow(false)
                                setSelected([])
                            }
                        }/>
                    </Box>
                </Drop>
            }
        </Box>
    )
}

function Pages({
    pages,
    current,
    onClickPage,
    ...props
}: {
    pages: (AlbumScreenAlbumPageCoverFragment | AlbumScreenAlbumPageVisualizationFragment)[],
    current: number,
    onClickPage: (n: number) => void
} & BoxExtendedProps) {
    return (
        <Box overflow="auto"     pad="xsmall" {...props}>
            <Grid
                columns="xsmall"
                style={{gridAutoFlow: "column", overflowX: "scroll"}}
                gap="xsmall"
                pad="xsmall"
            >
                {pages.map((p, key) => {
                    return (
                        <Box
                            key={key}
                            height="xsmall"
                            width="xsmall"
                            flex={{"shrink":0}}
                            style={{boxShadow: current === key ? "0 0 0px 2px #7D4CDB": "none" }}
                            align="center"
                        >
                            <Paper size="A4" scale={0.1}>
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