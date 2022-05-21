import React, { useEffect, useState } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text, Paragraph,  Layer } from "grommet"
import { StatusGood } from "grommet-icons"

import AnchorLink from "../common/AnchorLink"
import UserAvatar from "./UserAvatar/UserAvatar"

import { useAuthContext } from "../common/context/auth/useAuthContext"

import { useProject, Project  as ProjectType } from "./useProject"
import { ProjectFileType } from "./useUploadProjectFile"
import { ProjectEnums } from "../api/types"

import ChangeStatus from "./ChangeStatus/ChangeStatus"
import Contacts from "./Contacts/Contacts"
import Loading from "./Loading/Loading"
import AddSomething from "./AddSomething/AddSomething"
import ProjectDates from "./ProjectDates/ProjectDates"
import House from "./House/House"
import Rooms from "./Rooms/Rooms"
import Visualizations from "./Visualizations/Visualizations"
import UploadFiles from "./UploadFiles/UploadFiles"


interface RouteParams {
    id: string
};

// const spec = {
//     items: [
//         {
//             name: "PAX ПАКС / TYSSEDAL ТИССЕДАЛЬ Гардероб, комбинация",
//             image: "https://www.ikea.com/ru/ru/images/products/pax-paks-tyssedal-tissedal-garderob-kombinaciya-belyy__1024717_pe833642_s5.jpg?f=xl",
//             url: "https://www.ikea.com/ru/ru/p/pax-paks-tyssedal-tissedal-garderob-kombinaciya-belyy-s19429737/",
//         },
//         {
//             name: "MOALISA МОАЛИЗА Гардины, 2 шт. × 3",
//             image: "https://www.ikea.com/ru/ru/images/products/moalisa-moaliza-gardiny-2-sht-belyy-chernyy__0950019_pe801219_s5.jpg?f=xl",
//             url: "https://www.ikea.com/ru/ru/p/moalisa-moaliza-gardiny-2-sht-belyy-chernyy-50499515/",
//         },
//         {
//             name: "TYSSEDAL ТИССЕДАЛЬ Каркас кровати",
//             image: "https://www.ikea.com/ru/ru/images/products/tyssedal-tissedal-karkas-krovati-belyy-luroy__0637608_pe698420_s5.jpg?f=xl",
//             url: "https://www.ikea.com/ru/ru/p/tyssedal-tissedal-karkas-krovati-belyy-luroy-s79211165/"
//         },
//         {
//             name: "TYBBLE ТИБЛ Подвесной светильник/5 светодиодов",
//             image: "https://www.ikea.com/ru/ru/images/products/tybble-tibl-podvesnoy-svetilnik-5-svetodiodov-nikelirovannyy-molochnyy-steklo__0794591_pe765659_s5.jpg?f=xl",
//             url: "https://www.ikea.com/ru/ru/p/tybble-tibl-podvesnoy-svetilnik-5-svetodiodov-nikelirovannyy-molochnyy-steklo-90398251/",
//         }
//     ]
// };

export function Project () {
    let { id } = useParams<RouteParams>()

    const { user } = useAuthContext()
    const { data, loading, error } = useProject(id)

    const [ notification, setNotification ] = useState("")
    const [ showNotification, setShowNotification ] = useState(false)

    const notify = ({ message, timeout = 250, duration = 1500 }: { message: string, timeout?: number, duration?: number }) => {
        setNotification(message)
        
        setTimeout(() => {
            setShowNotification(true)

            setTimeout(() => {
                setShowNotification(false)
            }, duration)
        }, timeout)
    }

    const [ project, setProject ] = useState<ProjectType | undefined>(undefined)

    const [ projectEnums, setProjectEnums ] = useState<ProjectEnums | undefined>(undefined)

    useEffect(() => {
        switch (data?.screen.projectScreen.project.__typename) {
            case "Project":
                setProject(data?.screen.projectScreen.project)
                setProjectEnums(data?.screen.projectScreen.enums)
                return
        }
    }, [ data ])

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
        // default:
        //     return (
        //         <Main pad="large">
        //             <Heading>Ошибка</Heading>
        //             <Paragraph>Неизвестная ошибка</Paragraph>
        //         </Main>
        //     );
    }

    if (project) {
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
                        <AnchorLink to="/">apartomat</AnchorLink>
                    </Text>
                </Box>
                <Box><UserAvatar user={user} className="header-user" /></Box>
            </Header>

            <Box>
                <Box direction="row" justify="between" margin={{vertical: "medium"}}>
                    <Box direction="row" justify="center">
                        <Heading level={2} margin="none">{project.title}</Heading>
                        <ChangeStatus
                            projectId={project.id}
                            status={project.status}
                            values={projectEnums?.status}
                            onProjectStatusChanged={({ status }) => {
                                setProject({ ...project, status })
                            }}
                        />
                    </Box>
                    <AddSomething showUploadFiles={setShowUploadFiles}/>
                </Box>

                <Box direction="row" justify="between" wrap>
                    <Box width={{min: "35%"}}>
                        <Box margin="none">
                            <Heading level={4} margin={{ bottom: "xsmall"}}>Сроки проекта</Heading>
                            <ProjectDates
                                projectId={project.id}
                                startAt={project.startAt}
                                endAt={project.endAt}
                                onProjectDatesChanged={({ startAt, endAt }) => {
                                    notify({ message: "Даты изменены" })
                                    setProject({ ...project, startAt, endAt })
                                }}
                            />
                        </Box>
                        <Contacts contacts={project.contacts} projectId={project.id.toString()} notify={notify}/>
                    </Box>
                    <Box width={{min: "35%"}}>
                        <House houses={project.houses}/>
                        <Rooms houses={project.houses}/>
                    </Box>
                </Box>

                <Visualizations files={project.files} showUploadFiles={setShowUploadFiles}/>

                {/* <Box margin={{vertical: "medium"}}>
                    <Heading level={3}>
                        <Box gap="small">
                            <strong>Спецификация</strong>
                            <Text weight="normal">
                                {spec.items.length} позиций
                            </Text>
                        </Box>
                    </Heading>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}} >
                        {spec.items.map(item => {
                            return (
                                <Box
                                    key={item.name}
                                    height="xsmall"
                                    width="xsmall"
                                    margin={{bottom: "small"}}
                                    flex={{"shrink":0}}
                                    background="light-0"
                                >
                                    <Tip
                                        content={
                                            <Text>{item.name}</Text>
                                        }
                                    >
                                        <Image
                                            fit="cover"
                                            src={`${item.image}`}
                                        />
                                    </Tip>
                                </Box>
                            )
                        })}
                    </Box>
                </Box> */}

                {/* <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Планы</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Загрузить" />
                        </Box>
                    </Box>
                </Box>

                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Спецификация</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Создать" />
                        </Box>
                    </Box>
                </Box>

                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Альбом</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Создать" />
                        </Box>
                    </Box>
                </Box>

                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Исходники</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Загрузить" />
                        </Box>
                    </Box>
                </Box> */}

                {showUploadFiles ?
                    <UploadFiles
                        projectId={project.id}
                        type={ProjectFileType.Visualization}
                        setShow={setShowUploadFiles}
                        onUploadComplete={({message}) => notify({ message })}
                    /> : null}
                </Box>
            </Main>
        );
    } else {
        return <></>
    }
}

export default Project;