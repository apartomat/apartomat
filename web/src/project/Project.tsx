import React, { ChangeEvent, useState } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text,
    Paragraph, Spinner, SpinnerExtendedProps, Anchor, Image,
FileInput, RangeInput, Button } from "grommet"

import UserAvatar from "./UserAvatar";

import { useProject, ProjectFiles, MenuResult } from "./useProject"
import { useUploadProjectFile } from "./useUploadProjectFile"

import { useAuthContext } from "../common/context/auth/useAuthContext"
import { HeightType } from "grommet/utils";

interface RouteParams {
    id: string
};

export function Project () {
    let { id } = useParams<RouteParams>();

    const { user } = useAuthContext()
    const { data, loading, error } = useProject(parseInt(id, 10));

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
                        <Files files={project.files} />
                        <Upload projectId={project.id} />
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

function Files({ files }: { files: ProjectFiles }) {

    const [size, setSize] = useState(2)

    const onChange = (event: ChangeEvent<HTMLInputElement>) => setSize(parseInt(event.target.value, 10))

    console.log(size);

    const imageSize = (s: Number): HeightType => {
        switch (s) {
            case 1:
                return "xsmall"
            case 2:
                return "small"
            case 3:
                return "medium"
            default:
                return "small";
        }
    }

    switch (files.list.__typename) {
        case "ProjectFilesList":
            return (
                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Файлы</Heading>
                        <Box margin={{vertical: "medium", right: "medium"}} justify="center">
                            <RangeInput
                                min={1} max={3}step={1}
                                value={size}
                                onChange={onChange}
                            />
                        </Box>
                    </Box>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}} >
                        {files.list.items.map(file =>
                            <Box key={file.id} height={imageSize(size)} width={imageSize(size)} margin={{bottom: "small"}} flex={{"shrink":0}}>
                                <Image
                                    fit="cover"
                                    src={`${file.url}?w=192`}
                                    srcSet={`${file.url}?w=192 192w, ${file.url}?w=384 384w`}
                                />
                            </Box>
                        )}
                    </Box>
                </Box>
            )
        default:
            return <div>n/a</div>
    }
}

function Upload({ projectId }: {projectId: number}) {
    const [ file, setFile ] = useState<File | null>(null)
    const [ upload, { loading, error } ] = useUploadProjectFile()

    const onChange = (event: ChangeEvent<HTMLInputElement>) => {
        if (event.target.files) {
            setFile(event.target.files[0])
        }
    }

    const handleSubmit = (event: React.FormEvent) => {
        console.log({ file, loading, error });
        if (file && !loading) {
            upload({ projectId, file })
        }

        event.preventDefault();
    }

    return (
        <div>
            {error ? <p>{error.message}</p> : null}
            {loading ? null : <FileInput name="file" multiple onChange={onChange} messages={{browse:"выбрать", dropPrompt:"перетащите файл сюда", dropPromptMultiple:"перетащите файлы сюда"}}/>}
            {loading ? <p>Upload file...</p> : <Button disabled={!file} onClick={handleSubmit} label="Загрузить" />}
        </div>
    )
}

function Nav({ menu }: { menu: MenuResult }) {
    switch (menu.__typename) {
        case "MenuItems":
            const items = menu.items;

            return (
                <nav>
                    <ul style={{backgroundColor:'#eee'}}>
                        {
                            items.map(item => <li key={item.url}>
                                <a href={item.url}>{item.title}</a>
                            </li>)
                        }
                    </ul>
                </nav>
            )
    }
    
    return null;
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