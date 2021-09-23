import React, { ChangeEvent, Dispatch, SetStateAction, useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text,
    Paragraph, Spinner, SpinnerExtendedProps, Anchor, Image,
FileInput, Button, Layer, Form, FormField } from "grommet"
import { FormClose, StatusGood } from 'grommet-icons'

import UserAvatar from "./UserAvatar"

import { useProject, ProjectFiles } from "./useProject"
import { useUploadProjectFile } from "./useUploadProjectFile"

import { useAuthContext } from "../common/context/auth/useAuthContext"

interface RouteParams {
    id: string
};

export function Project () {
    let { id } = useParams<RouteParams>();

    const { user } = useAuthContext()
    const { data, loading, error } = useProject(parseInt(id, 10));

    const [ notification, setNotification ] = useState('')
    const [ showNotification, setShowNotification ] = useState(false)

    const notify = ({ message }: { message: string }) => {
        setNotification(message)
        setShowNotification(true)
        setTimeout(() => {
            setShowNotification(false)
        }, 3000)
    }

    const [showUploadFiles, setShowUploadFiles] = useState(false);

    if (loading) {
        return (
            <Main pad="large">
                <Box direction="row" gap="small" align="center">
                    <Loading message="Загрузка..."/>
                    <Text>Загрузка...</Text>
                </Box>
            </Main>
        );
    }

    if (error) {
        return (
            <Main pad="large">
                <Heading>Ошибка</Heading>
                <Paragraph>{error}</Paragraph>
            </Main>
        );
    }

    switch (data?.screen.projectScreen.project.__typename) {
        case "Project":
            const { project, menu } = data?.screen.projectScreen;

            return (
                <Main pad={{vertical: "medium", horizontal: "large"}}>

                    {showNotification ? <Layer
                        position="top"
                        modal={false}
                        responsive={false}
                        margin={{ vertical: "small", horizontal: "small"}}
                    >
                        <Box
                            align="center"
                            direction="row"
                            gap="xsmall"
                            justify="between"
                            elevation="small"
                            background="status-ok"
                            round="medium"
                            pad={{ vertical: "xsmall", horizontal: "small"}}
                        >
                            <StatusGood/>
                            <Text>{notification}</Text>
                        </Box>
                    </Layer> : null}

                    <Header background="white" margin={{vertical: "medium"}}>
                        <Box>
                            
                            <Text size="xlarge" weight="bold" color="brand">
                                <Anchor href="/">apartomat</Anchor>
                            </Text>
                        </Box>
                        <Box><UserAvatar user={user} className="header-user" /></Box>
                    </Header>

                    <Box>
                        <Box margin={{vertical: "medium"}}>
                            <Heading level={2} margin="none">{project.title}</Heading>
                        </Box>
                        <Box direction="row" justify="between" wrap={true}>
                            <Box width="medium">
                                <Box margin="none">
                                    <Heading level={4} margin={{bottom: "xxsmall"}}>Сроки проекта</Heading>
                                    <Text margin={{top: "xxsmall"}}>2021/08/12&mdash;2021/09/12</Text>
                                </Box>
                                <Box margin={{top: "small"}}>
                                    <Heading level={4} margin={{bottom: "xxsmall"}}>Комплектация</Heading>
                                    <Text margin={{top: "xxsmall"}}>3 комнаты, 2 санузла, коридор</Text>
                                </Box>
                            </Box>
                            <Box width="medium">
                                <Box margin="none">
                                    <Heading level={4} margin={{bottom: "xxsmall"}}>Адрес</Heading>
                                    <Text margin={{top: "xxsmall"}}>Москва, ул. Амурская, ЖК LEVEL</Text>
                                </Box>
                                <Box margin={{top: "small"}}>
                                    <Heading level={4} margin={{bottom: "xxsmall"}}>Заказчик</Heading>
                                    <Text margin={{top: "xxsmall"}}><Anchor href="">Екатерина</Anchor>, <Anchor href="">Иван</Anchor></Text>
                                </Box>
                            </Box>
                        </Box>

                        <Files files={project.files} showUploadFiles={setShowUploadFiles}/>

                        <Box>
                            <Box direction="row" justify="between">
                                <Heading level={3}>Спецификация</Heading>
                            </Box>
                        </Box>

                        <Box>
                            <Box direction="row" justify="between">
                                <Heading level={3}>Альбом</Heading>
                            </Box>
                        </Box>

                        {showUploadFiles ?
                            <UploadFiles
                                projectId={project.id}
                                setShow={setShowUploadFiles}
                                onUploadComplete={({message}) => notify({ message })}
                            /> : null}
                    </Box>
                </Main>
            );
        case "NotFound":
            return (
                <Main pad="large">
                    <Heading level={2}>Ошибка</Heading>
                    <Box>
                        <Text>Проект не найден</Text>
                    </Box>
                </Main>
            );
        case "Forbidden":
            return (
                <Main pad="large">
                    <Heading level={2}>Ошибка</Heading>
                    <Paragraph>Доступ запрещен</Paragraph>
                </Main>
            );
        default:
            return (
                <Main pad="large">
                    <Heading>Ошибка</Heading>
                    <Paragraph>Неизвестная ошибка</Paragraph>
                </Main>
            );
    }
}

export default Project;

function Files({ files, showUploadFiles }: { files: ProjectFiles, showUploadFiles: Dispatch<SetStateAction<boolean>> }) {

    const handleUploadFiles = () => {
        showUploadFiles(true)
    }

    switch (files.list.__typename) {
        case "ProjectFilesList":
            return (
                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Файлы</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Загрузить файлы" onClick={handleUploadFiles} />
                        </Box>
                    </Box>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}} >
                        {files.list.items.map(file =>
                            <Box key={file.id} height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}}>
                                <Image
                                    fit="cover"
                                    src={`${file.url}?w=192`}
                                    srcSet={`${file.url}?w=192 192w, ${file.url}?w=384 384w`}
                                />
                            </Box>
                        )}
                        {files.total.__typename === 'ProjectFilesTotal' && files.total.total > files.list.items.length
                            ? <Box height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}} align="center" justify="center">
                                <Text>ещё файлов {files.total.total - files.list.items.length}</Text>
                            </Box>
                            : null
                        }
                    </Box>
                </Box>
            )
        default:
            return <div>n/a</div>
    }
}

function UploadFiles(
    { projectId, setShow, onUploadComplete }:
    {
        projectId: number,
        setShow: Dispatch<SetStateAction<boolean>>,
        onUploadComplete: ({message}: { message: string}) => void
    }
) {
    const [ files, setFiles ] = useState<FileList | null>(null)
    const [ upload, { loading, error, data, called } ] = useUploadProjectFile()

    console.log({ loading, error, data, called });

    const handleSubmit = (event: React.FormEvent) => {
        if (files && !loading) {
            upload({ projectId, file: files[0] })
        }

        event.preventDefault()
    }

    const handleFileInputChange = (event: ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            setFiles(event.target.files)
        }
    }

    useEffect(() => {
        if (called && data && !error) {
            setShow(false)
            onUploadComplete({ message: files?.length === 1 ? "Файл загружен" : `Загружено файлов ${files?.length}`})
        }
    })

    return (
        <Layer
            onClickOutside={() => setShow(false)}
            onEsc={() => setShow(false)}
        >
            <Form validate="submit" onSubmit={handleSubmit}>
                <Box pad="medium" gap="medium" width="large">
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Загрузить файлы</Heading>
                        <Button icon={ <FormClose/> } onClick={() => setShow(false)}/>
                    </Box>
                    <FormField name="files" htmlFor="files" required margin="none">
                        <FileInput
                            name="files"
                            renderFile={(file) => (
                                <Box pad="small" direction="row" justify="between" align="center">
                                    <Box width="xsmall" height="xsmall">
                                        <Image src={ URL.createObjectURL(file) } fit="cover" />
                                    </Box>
                                    <Text>{file.name}</Text>
                                </Box>
                            )}
                            multiple={{"aggregateThreshold": 5}}
                            messages={{
                                browse: "выбрать",
                                dropPrompt: "перетащите файл сюда",
                                dropPromptMultiple: "перетащите файлы сюда",
                                files: "файлов",
                                remove: "удалить",
                                removeAll: "удалить все"
                            }}
                            onChange={handleFileInputChange}
                        />
                    </FormField>
                    <Box align="center">
                        <Button
                            type="submit"
                            primary
                            label={loading ? 'Загрузка...' : 'Загрузить' }
                            disabled={loading}
                        />
                    </Box>
                </Box>
            </Form>
        </Layer>
    )
}

const Loading = (props: SpinnerExtendedProps) => {
    return (
        <Spinner
            border={[
                { side: 'all', color: 'background-contrast', size: 'medium' },
                { side: 'right', color: 'brand', size: 'medium' },
                { side: 'top', color: 'brand', size: 'medium' },
                { side: 'left', color: 'brand', size: 'medium' },
            ]}
            {...props}
        />
    )
}