import { Box, Button, Grid, Heading, Image, Main, Text, Header } from "grommet"
import { DocumentPdf } from "grommet-icons"
import { memo, useEffect, useState } from "react"
import Lightbox from "yet-another-react-lightbox"
import "yet-another-react-lightbox/styles.css"
import { useProjectPage, ProjectPage as ProjectPageType, House, Album } from "./api"
import { filesize } from "filesize"

export function ProjectPage({ id }: { id: string }) {
    const { data, loading } = useProjectPage(id)

    const [projectPage, setProjectPage] = useState<ProjectPageType>()

    const [error, setError] = useState<string | undefined>()

    const [house, setHouse] = useState<House | undefined>()

    const [visualizations, setVisualizations] = useState([])

    const [album, setAlbum] = useState<Album | undefined>()

    useEffect(() => {
        const res = data?.projectPage

        if (!res) {
            return
        }

        switch (res?.__typename) {
            case "ProjectPage":
                setProjectPage(res)

                if (res.house.__typename === "House") {
                    setHouse(res.house)
                }

                if (res.visualizations.list.__typename === "VisualizationsList") {
                    setVisualizations(res.visualizations.list.items)
                }

                if (res.album.__typename === "Album") {
                    setAlbum(res.album)
                }

                break
            case "NotFound":
                setError("Проект не найден")
                break
            case "Forbidden":
                setError("Доступ запрещен")
                break
            default:
                setError("Ошибка сервера")
                break
        }
    }, [data])

    const [showAlbumFullscreen, setShowAlbumFullscreen] = useState<number | undefined>()

    if (loading) {
        return (
            <Main pad="large">
                <Box direction="row" gap="small" align="center">
                    <Text>Загрузка...</Text>
                </Box>
            </Main>
        )
    }

    if (error) {
        return (
            <Main pad="large">
                <Heading level={2}>Ошибка</Heading>
                <Box>
                    <Text>{error}</Text>
                </Box>
            </Main>
        )
    }

    if (!projectPage) {
        return <></>
    }

    return (
        <Box pad={{ vertical: "medium", horizontal: "large" }}>
            <Header margin={{ vertical: "medium" }}>
                <Box gap="small">
                    <Heading level={2} margin="none" color="brand">
                        {projectPage.title}
                    </Heading>
                    <Box>
                        <Text size="small">{projectPage.description}</Text>
                    </Box>
                </Box>
            </Header>

            <Main>
                {house && (
                    <Box pad={{ horizontal: "xxsmall", vertical: "small" }}>
                        <Box direction="row" gap="small">
                            <Address house={house} />
                        </Box>
                    </Box>
                )}

                {visualizations.length > 0 && (
                    <Box margin={{ bottom: "large" }}>
                        <Box direction="row">
                            <Heading level={3}>Визуализации</Heading>
                        </Box>
                        <Box overflow="auto">
                            <Grid columns="small" style={{ gridAutoFlow: "column", overflowX: "scroll" }} gap="xsmall">
                                {visualizations.map((vis, index) => (
                                    <Box width="small" height="small" background="light-2">
                                        <Image
                                            src={vis.file.url}
                                            fit="contain"
                                            onDoubleClick={() => setShowAlbumFullscreen(index)}
                                        />
                                    </Box>
                                ))}
                            </Grid>
                        </Box>
                    </Box>
                )}

                {album && (
                    <Box margin={{ bottom: "large" }}>
                        <Box direction="row" justify="between">
                            <Heading level={3}>Файлы</Heading>
                        </Box>
                        {album && (
                            <Box direction="row" gap="small" align="center">
                                <Heading level="5" margin="none">
                                    Альбом визуализаций
                                </Heading>
                                <Button label={"Скачать"} icon={<DocumentPdf />} href={album.url} />
                                <Box>
                                    <Text size="small">{albumFilesize(album)}</Text>
                                </Box>
                            </Box>
                        )}
                    </Box>
                )}

                <Lightbox
                    open={showAlbumFullscreen !== undefined}
                    close={() => setShowAlbumFullscreen(undefined)}
                    slides={visualizations.map((vis) => {
                        return { src: vis.file.url }
                    })}
                    index={showAlbumFullscreen || 0}
                ></Lightbox>
            </Main>
        </Box>
    )
}

const Address = memo(function Address({ house: { city, housingComplex, address } }: { house: House }) {
    const short = []
    const full = []

    const sep = ", "

    if (city) {
        short.push(city)
        full.push(city)
    }

    if (housingComplex) {
        short.push(`ЖК «${housingComplex}»`)
        full.push(`ЖК «${housingComplex}»`)
    }

    if (address) {
        full.push(address)
    }

    return <Button primary color="light-2" label={short.join(sep)} title={full.join(sep)} />
})

function albumFilesize(album: Album) {
    return filesize(album.size, { locale: "ru" })
}
