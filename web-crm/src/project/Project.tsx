import React, { ChangeEvent, Dispatch, SetStateAction, useEffect, useState, useRef, useMemo } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text,
    Paragraph, Spinner, SpinnerExtendedProps, Image,
FileInput, Button, Layer, Form, FormField, Drop, MaskedInput, Card, CardHeader, CardBody, CardFooter, Calendar } from "grommet"
import { FormClose, StatusGood, Add, Trash, Instagram, Next } from "grommet-icons"

import AnchorLink from "../common/AnchorLink"
import UserAvatar from "./UserAvatar"


import { useAuthContext } from "../common/context/auth/useAuthContext"

import { useProject, Project  as ProjectType, ProjectFiles, ProjectContacts, ProjectHouses, HouseRooms, ProjectStatus } from "./useProject"
import { useUploadProjectFile, ProjectFileType } from "./useUploadProjectFile"
import { useAddContact, ContactType, Contact } from "./useAddContact"
import { useUpdateContact } from "./useUpdateContact"
import { useDeleteContact } from "./useDeleteContact"
import { useChangeDates } from "./useChangeDates"
import { ProjectEnums, ProjectStatusEnum, ProjectStatusEnumItem } from "../api/types"
import { useChangeStatus } from "./useChangeStatus"

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
                        <ChangeProjectStatus
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

function statusToLabel({ status, items }: { status: ProjectStatus, items?: ProjectStatusEnumItem[] }): string {
    if (!items) {
        return ""
    }

    for (let item of items) {
        if (item.key === status) {
            return item.value
        }
    }

    return ""
}

function statusColor(status: ProjectStatus): string {
    switch(status) {
        case ProjectStatus.New:
            return "status-unknown"
        case ProjectStatus.InProgress:
            return "status-ok"
        case ProjectStatus.Done:
            return "status-warning"
        case ProjectStatus.Canceled:
            return "status-error"
    }
}

function ChangeProjectStatus({
    projectId,
    status,
    values,
    onProjectStatusChanged
}: {
    projectId: string,
    status: ProjectStatus,
    values?: ProjectStatusEnum,
    onProjectStatusChanged?: ({ status }: { status: ProjectStatus }) => void
}) {
    const [ show, setShow ] = useState<Boolean>(false)
    
    const [ state, setState ] = useState(status)

    const [ changeStatus, { data, loading, error }] = useChangeStatus()

    const handleItemClick = (projectId: string, status: ProjectStatus) => {
        changeStatus(projectId, status)
        setShow(false)
        setState(status)
    }

    useEffect(() => {
        switch (data?.changeProjectStatus.__typename) {
            case "ProjectStatusChanged":
                const { status } = data?.changeProjectStatus.project
                onProjectStatusChanged && onProjectStatusChanged({ status })
        }
    }, [ data, onProjectStatusChanged ])

    useEffect(() => {
        if (error) {
            setTimeout(() => setState(status), 300)
        }
    })

    const label = useMemo(() => statusToLabel({ status: state, items: values?.items }), [ state, values ])

    const color = useMemo(() => statusColor(state), [ state ])

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box justify="center" margin={{ horizontal: "medium"}}>
            <Box ref={targetRef}>
                <Button
                    label={label}
                    color={color}
                    size="small"
                    onClick={() => setShow(true)}
                    disabled={loading}
                />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{ top: "bottom", left: "left" }}
                    margin={{ top: "xsmall" }}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                >
                    {values?.items.map(item => {
                        return (
                            <Button key={item.key} plain hoverIndicator={{color: "light-2"}}>
                                <Box pad="small" onClick={() => handleItemClick(projectId, item.key)}><Text>{item.value}</Text></Box>
                            </Button>
                        )
                    })}
                </Drop>
            )}
        </Box>
    )
}

function ContactCard(
    { contact, onDelete, onClickUpdate }:
    { contact: Contact , onDelete: (contact: Contact) => void, onClickUpdate: (contact: Contact) => void }
) {
    const ref = useRef(null)

    const [showCard, setShowCard] = useState(false)

    const [showDeleteConfirm, setShowDeleteConfirm] = useState(false)

    const [deleteContact, { data } ] = useDeleteContact()

    const handleDelete = () => {
        setShowDeleteConfirm(true)
    }

    const handleDeleteConfirm = () => {
        deleteContact(contact.id)
    }

    const handleDeleteCancel = () => {
        setShowDeleteConfirm(false)
        setShowCard(false)
    }

    useEffect(() => {
        switch (data?.deleteContact.__typename) {
            case "ContactDeleted":
                setShowCard(false)
                onDelete(contact)
        }
    }, [ data, contact ]) //todo

    return (
        <Box>
            <Button
                key={contact.id}
                ref={ref}
                primary
                color="light-2"
                label={contact.fullName}
                style={{whiteSpace: "nowrap"}}
                onClick={() => setShowCard(!showCard) }
            />
            {ref.current && showCard &&
                <Drop
                    target={ref.current}
                    align={{left: "right"}}
                    plain
                    onEsc={() => setShowCard(false) }
                    onClickOutside={() => setShowCard(false) }
                >
                    <Card width="medium" background="white" margin="small">
                        <CardHeader pad={{horizontal: "medium", top: "medium"}} style={{fontWeight: "bold"}}>{contact.fullName}</CardHeader>
                        <CardBody pad="medium">
                            {contact.details.filter(c => ![ContactType.Instagram].includes(c.type)).map(c => {
                                return <Box pad={{vertical: "small"}}>{c.value}</Box>
                            })}
                            {contact.details.filter(c => [ContactType.Instagram].includes(c.type)).length > 0
                                ? <Box pad={{vertical: "small"}} direction="row">
                                {contact.details.filter(c => [ContactType.Instagram].includes(c.type)).map(c => {
                                    switch (c.type) {
                                        case ContactType.Instagram:
                                            return <Button icon={<Instagram color="primary"/>} plain href={c.value}/>
                                    }

                                    return null
                                })}
                                </Box>
                                : null
                            }
                        </CardBody>
                        <CardFooter pad={{horizontal: "small"}} background="light-1" height="xxsmall">
                            {showDeleteConfirm
                                ?
                                    <Box direction="row" gap="small">
                                        <Button primary label="Удалить" size="small" onClick={handleDeleteConfirm}/>
                                        <Button label="Отмена" size="small" onClick={handleDeleteCancel}/>
                                    </Box>
                                : (
                                    <Button icon={<Trash/>} onClick={handleDelete}/>
                                )
                            }

                            {!showDeleteConfirm &&
                                <Button label="Редактировать" size="small" primary onClick={() => {
                                    setShowCard(false)
                                    onClickUpdate(contact)
                                }}/>
                            }

                        </CardFooter>
                    </Card>
                </Drop>
            }
        </Box>
    )
}

function Contacts({ contacts, projectId, notify }: { contacts: ProjectContacts, projectId: string, notify: (val: { message: string }) => void }) {
    const [showAddContact, setShowAddContact] = useState(false)

    const [updateContact, setUpdateContact] = useState<Contact | undefined>(undefined)

    const [justAdded, setJustAdded] = useState([] as Contact[])

    const [justDeleted, setJustDeleted] = useState([] as string[])

    switch (contacts.list.__typename) {
        case "ProjectContactsList":

            const list = [...contacts.list.items, ...justAdded].filter(contact => !justDeleted.includes(contact.id))

            return (
                <>
                    <Box margin={{top: "small"}}>
                        <Heading level={4} margin={{ bottom: "xsmall"}}>
                            {list.length === 1 ? "Заказчик" : "Заказчики"}
                        </Heading>
                        <Box direction="row" gap="small" wrap>
                            {[...list.map((contact) => {
                                return (
                                    <ContactCard
                                        key={contact.id}
                                        contact={contact}
                                        onDelete={(contact: Contact) => {
                                            setJustDeleted([...justDeleted, contact.id])
                                            notify({ message: "Контакт удален"})
                                        }}
                                        onClickUpdate={(contact: Contact) => {
                                            setUpdateContact(contact)
                                        }}
                                    />
                                )
                            }), <Button key="" icon={<Add/>} label="Добавить" onClick={() => setShowAddContact(true) }/>]}
                        </Box>
                    </Box>

                    {showAddContact ?
                        <AddContact
                            projectId={projectId}
                            setShow={setShowAddContact}
                            onAdd={(contact: Contact) => {
                                setJustAdded([...justAdded, contact])
                                notify({ message: "Контакт добавлен"})
                            }}
                        /> : null}

                    {updateContact ?
                        <UpdateContact
                            contact={updateContact}
                            hide={() => { setUpdateContact(undefined) }}
                            onUpdate={(contact: Contact) => {
                                notify({ message: "Контакт сохранен"})
                            }}
                        /> : null}
                </>
            )
        default:
            return (
                <Box margin={{top: "small"}}>
                    <Text>n/a</Text>
                </Box>
            )
    }
}

function HouseText({ houses }: { houses: ProjectHouses }) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            const [ house ] =  houses.list.items

            if (!house) {
                return (
                    <>n/a</>
                )
            }

            return (
                <>{[house.city, house.address, house.housingComplex].join(', ')}</>
            )
        default:
            return (
                <>n/a</>
            )
    }
}

function House({ houses }: { houses: ProjectHouses }) {
    return (
        <Box margin="none">
            <Heading level={4} margin={{ bottom: "xsmall"}}>Адрес</Heading>
            <Box height="36px" justify="center">
                <Text><HouseText houses={houses}/></Text>
            </Box>
        </Box>
    )
}


function RoomsText({ rooms }: { rooms: HouseRooms }) {
    switch (rooms.list.__typename) {
        case "HouseRoomsList":
            if (rooms.list.items.length === 0) {
                return (
                    <Text>n/a</Text>
                )
            }

            return (
                <Text>{rooms.list.items.length} помещений, {rooms.list.items.reduce((acc, room) => {
                    return acc + (room.square || 0)
                }, 0)} м<sup>2</sup></Text>
            )
        default:
            return (
                <Text>n/a</Text>
            )
    }
}

function Rooms({ houses }: { houses: ProjectHouses }) {
    const house = firstHouse(houses)

    return (
        <Box margin={{top: "small"}}>
            <Heading level={4} margin={{ bottom: "xsmall"}}>Комплектация</Heading>
            <Box height="36px" justify="center">
                {
                    (house ? <RoomsText rooms={house.rooms}/> : <Text>n/a</Text>)
                }
            </Box>
        </Box>
    )
}

function firstHouse(houses: ProjectHouses) {
    switch (houses.list.__typename) {
        case "ProjectHousesList":
            return houses.list.items[0]
        default:
            return undefined
    }
}

function ProjectDates ({
    projectId,
    startAt,
    endAt,
    onProjectDatesChanged
}: {
    projectId: string,
    startAt?: string,
    endAt?: string,
    onProjectDatesChanged?: (dates: { startAt?: string, endAt?: string }) => void,
}) {
    const [ showChangeDates, setShowChangeDates ] = useState(false)

    const [ label, setLabel ] = useState(<>не определены</>)

    useEffect(() => {
        if (startAt && endAt) {
            return setLabel(
                <>
                    {new Date(startAt).toLocaleDateString("ru-RU")}&nbsp;&mdash;&nbsp;{new Date(endAt).toLocaleDateString("ru-RU")}
                </>
            )
        }

        if (startAt) {
            return setLabel(<>new Date(startAt).toLocaleDateString("ru-RU")</>)
        }

        return setLabel(<>не определены</>)
    }, [ startAt, endAt ])

    return (
        <>
            <Box direction="row">
                <Button
                    primary
                    color="light-2"
                    label={label}
                    onClick={() => setShowChangeDates(!showChangeDates)}
                />
            </Box>
            {showChangeDates &&
                <ChangeDates
                    projectId={projectId}
                    startAt={startAt}
                    endAt={endAt}
                    onEsc={() => setShowChangeDates(false) }
                    onClickOutside={() => setShowChangeDates(false) }
                    onClickClose={() => setShowChangeDates(false) }
                    onProjectDatesChanged={({ startAt, endAt }) => {
                        onProjectDatesChanged && onProjectDatesChanged({ startAt, endAt })
                        setShowChangeDates(false)
                    }}
                />
            }
        </>
    )
}

function AddSomething ({ showUploadFiles }: { showUploadFiles: Dispatch<SetStateAction<boolean>> }) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box justify="center">
            <Box ref={targetRef}>
                <Button label="Добавить" icon={<Next/>} reverse onClick={() => setShow(true)} />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{top: "bottom", right: "right"}}
                    margin={{top: "xsmall"}}
                    target={targetRef.current}
                    onClickOutside={() => setShow(false)}
                    onEsc={() => setShow(false)}
                >
                    <Button plain children={({ hover }: {hover: boolean}) => {
                        return <Box pad="small" background={hover ? 'light-1': ''}><Text>Визуализации</Text></Box>
                    }} onClick={() => {
                        showUploadFiles(true)
                        setShow(false)
                    }}/>
                    <Button plain>
                        <Box pad="small"><Text>План</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Исходники</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Альбом</Text></Box>
                    </Button>
                    <Button plain>
                        <Box pad="small"><Text>Спецификация</Text></Box>
                    </Button>
                </Drop>
            )}
        </Box>
    )
}

function Visualizations({ files, showUploadFiles }: { files: ProjectFiles, showUploadFiles: Dispatch<SetStateAction<boolean>> }) {

    const handleUploadFiles = () => {
        showUploadFiles(true)
    }

    switch (files.list.__typename) {
        case "ProjectFilesList":
            if (files.list.items.length === 0) {
                return null
            }

            return (
                <Box margin={{vertical: "medium"}}>
                    <Box direction="row" justify="between">
                        <Heading level={3}>Визуализации</Heading>
                        <Box justify="center">
                            <Button color="brand" label="Загрузить" onClick={handleUploadFiles} />
                        </Box>
                    </Box>
                    <Box direction="row" gap="small" overflow={{"horizontal":"auto"}} >
                        {files.list.items.map(file =>
                            <Box
                                key={file.id}
                                height="small"
                                width="small"
                                margin={{bottom: "small"}}
                                flex={{"shrink":0}}
                                background="light-2"
                            >
                                <Image
                                    fit="cover"
                                    src={`${file.url}?w=192`}
                                    srcSet={`${file.url}?w=192 192w, ${file.url}?w=384 384w`}
                                />
                            </Box>
                        )}
                        {files.total.__typename === 'ProjectFilesTotal' && files.total.total > files.list.items.length
                            ? <Box key={0} height="small" width="small" margin={{bottom: "small"}} flex={{"shrink":0}} align="center" justify="center">
                                <Text>ещё {files.total.total - files.list.items.length}</Text>
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
    { projectId, type, setShow, onUploadComplete }:
    {
        projectId: string,
        type: ProjectFileType,
        setShow: Dispatch<SetStateAction<boolean>>,
        onUploadComplete: ({message}: { message: string}) => void
    }
) {
    const [ files, setFiles ] = useState<FileList | null>(null)
    const [ upload, { loading, error, data, called }, state ] = useUploadProjectFile()
    
    useEffect(() => {
        if (state.state === "DONE") {
            console.log(state.file);
        }

        if (state.state === "FAILED") {
            if (state.error instanceof Error) {
                console.log("------------", state.error.message)
            } else {
                console.log(state.error.__typename, state.error.message)
            }
            
        }
    })

    const handleSubmit = (event: React.FormEvent) => {
        if (files && !loading) {
            upload({ projectId, file: files[0], type })
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

function AddContact(
    { projectId, setShow, onAdd }:
    {
        projectId: string,
        setShow: Dispatch<SetStateAction<boolean>>,
        onAdd: (contact: Contact) => void
    }) {

    const [ value, setValue ] = useState({} as ContactFormData)

    const [ add, { data, loading, error } ] = useAddContact()

    const handleSubmit = (event: React.FormEvent) => {
        const { fullName } = value;

        const details = [];

        if (value.phone) {
            details.push({type: ContactType.Phone, value: value.phone});
        }

        if (value.email) {
            details.push({type: ContactType.Email, value: value.email});
        }

        if (value.instagram) {
            details.push({type: ContactType.Instagram, value: value.instagram});
        }

        add(projectId, { fullName, details })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.addContact.__typename) {
            case "ContactAdded":
                const { addContact: { contact }} = data
                onAdd(contact)
                setShow(false)
        }
    }, [ data, setShow, onAdd ]) // todo

    return (
        <Layer
            onClickOutside={() => setShow(false)}
            onEsc={() => setShow(false)}
        >
                {error && <Box><Text>{error.message}</Text></Box>}

                <Box pad="medium" gap="medium" width={{min: "500px"}}>
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Добавить контакт</Heading>
                        <Button icon={ <FormClose/> } onClick={() => setShow(false)}/>
                    </Box>
                    <ContactForm
                        contact={value}
                        onSet={setValue}
                        onSubmit={handleSubmit}
                        submit={
                            <Box direction="row" justify="between" margin={{top: "large"}}>
                                <Button
                                    type="submit"
                                    primary
                                    label={loading ? 'Сохранение...' : 'Сохранить' }
                                    disabled={loading}
                                />
                                <Box><Text color="status-critical"><ErrorMessage res={data?.addContact}/></Text></Box>
                            </Box>
                        }
                    />
                </Box>
        </Layer>
    )
}

type ContactFormData = { fullName: string, phone: string, email: string, instagram: string }

function ContactForm({
    contact,
    onSet,
    onSubmit,
    submit
}: {
    contact: ContactFormData,
    onSet: React.Dispatch<ContactFormData>,
    onSubmit: (event: React.FormEvent) => void,
    submit: JSX.Element
}) {
    return (
        <Form
            validate="submit"
            value={contact}
            onChange={val => onSet(val)}
            onSubmit={onSubmit}
            messages={{required: 'обязательное поле'}}
        >
            <FormField
                label="Имя"
                name="fullName"
                validate={{
                    regexp: /^.+$/,
                    message: "обязательно для заполнения",
                    status: "error"
                }}
            >
                <MaskedInput
                    name="fullName"
                    mask={[
                        { regexp: /^.*$/, placeholder: "Имя" },
                        { fixed: ' ' },
                        { regexp: /^.*$/, placeholder: "Фамилия" },
                    ]}
                />
            </FormField>
            <FormField label="Телефон" name="phone">
                <MaskedInput
                    name="phone"
                    mask={[
                        { fixed: '+7 (' },
                        {
                            length: 3,
                            regexp: /^[0-9]{1,3}$/,
                            placeholder: 'xxx',
                        },
                        { fixed: ')' },
                        { fixed: ' ' },
                        {
                            length: 3,
                            regexp: /^[0-9]{1,3}$/,
                            placeholder: 'xxx',
                        },
                        { fixed: '-' },
                        {
                            length: 2,
                            regexp: /^[0-9]{1,4}$/,
                            placeholder: 'xx',
                        },
                        { fixed: '-' },
                        {
                            length: 2,
                            regexp: /^[0-9]{1,4}$/,
                            placeholder: 'xx',
                        },
                    ]}
                />
            </FormField>
            <FormField label="Электронная почта" name="email">
                <MaskedInput
                    name="email"
                    mask={[
                        { regexp: /^[\w\-_.]+$/, placeholder: "example" },
                        { fixed: '@' },
                        { regexp: /^[\w\-_.]+$/, placeholder: "test" },
                        { fixed: '.' },
                        { regexp: /^[\w]+$/, placeholder: 'org' },
                    ]}
                />
            </FormField>
            <FormField label="Instagram" name="instagram">
                <MaskedInput
                    name="instagram"
                    mask={[{ fixed: 'https://www.instagram.com/' }, { regexp: /^.*$/ }]}
                />
            </FormField>
            {submit}
        </Form>
    )
}

function ChangeDates({
    projectId,
    startAt,
    endAt,
    onProjectDatesChanged,
    onEsc,
    onClickOutside,
    onClickClose
}: {
    projectId: string,
    startAt?: string,
    endAt?: string,
    onProjectDatesChanged?: (dates: { startAt?: string, endAt?: string }) => void,
    onEsc?: () => void,
    onClickOutside?: () => void,
    onClickClose?: () => void
}) {
    const [ dates, setDates ] = useState(startAt && endAt ? [[startAt, endAt]] : undefined)

    const [ change, { loading, data } ] = useChangeDates()

    useEffect(() => {
        switch (data?.changeProjectDates.__typename) {
            case "ProjectDatesChanged":
                onProjectDatesChanged && onProjectDatesChanged({ startAt: dates && dates[0] && dates[0][0], endAt: dates && dates[0] && dates[0][1] })
        }
    }, [ data, dates, onProjectDatesChanged ])


    const handleSelect = (value: any) => {
        setDates(value)
    }

    const handleSubmit = (event: React.FormEvent) => {
        change(projectId, { startAt: dates && dates[0] && dates[0][0], endAt: dates && dates[0] && dates[0][1] })
        event.preventDefault()
    }

    return (
        <Layer
            onClickOutside={onClickOutside}
            onEsc={onEsc}
        >
            <Box pad="medium" gap="medium">
                <Box direction="row" justify="between" align="center">
                    <Heading level={2} margin="none">Сроки проекта</Heading>
                    <Button icon={ <FormClose/> } onClick={onClickClose}/>
                </Box>
                <Box>
                    <Calendar
                        firstDayOfWeek={1}
                        locale="ru-RU"
                        range="array"
                        activeDate={undefined}
                        dates={dates}
                        onSelect={handleSelect}
                    />
                </Box>
                <Box direction="row" margin={{top: "medium"}}>
                    <Button type="submit" primary label="Сохранить" onClick={handleSubmit} disabled={loading}/>
                </Box>
            </Box>
        </Layer>
    )
}

function UpdateContact(
    { contact, hide, onUpdate }:
    {
        contact: Contact,
        hide: () => void,
        onUpdate: (contact: Contact) => void
    }) {

    const [ value, setValue ] = useState({
        fullName: contact.fullName,
        phone: contact.details.filter(val => val.type === ContactType.Phone)[0]?.value,
        email: contact.details.filter(val => val.type === ContactType.Email)[0]?.value,
        instagram: contact.details.filter(val => val.type === ContactType.Instagram)[0]?.value
    } as ContactFormData)

    const [ update, { data, loading, error } ] = useUpdateContact()

    const handleSubmit = (event: React.FormEvent) => {
        const { fullName } = value;

        const details = [];

        if (value.phone) {
            details.push({type: ContactType.Phone, value: value.phone});
        }

        if (value.email) {
            details.push({type: ContactType.Email, value: value.email});
        }

        if (value.instagram) {
            details.push({type: ContactType.Instagram, value: value.instagram});
        }

        update(contact.id, { fullName, details })

        event.preventDefault()
    }

    useEffect(() => {
        switch (data?.updateContact.__typename) {
            case "ContactUpdated":
                hide()
                onUpdate(data.updateContact.contact)
        }
    }, [ data, hide, onUpdate ])

    return (
        <Layer
            onClickOutside={hide}
            onEsc={hide}
        >
                {error && <Box><Text>{error.message}</Text></Box>}

                <Box pad="medium" gap="medium" width={{min: "500px"}}>
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Изменить контакт</Heading>
                        <Button icon={ <FormClose/> } onClick={hide}/>
                    </Box>
                    <ContactForm
                        contact={value}
                        onSet={setValue}
                        onSubmit={handleSubmit}
                        submit={
                            <Box direction="row" justify="between" margin={{top: "large"}}>
                                <Button
                                    type="submit"
                                    primary
                                    label={loading ? 'Сохранение...' : 'Сохранить' }
                                    disabled={loading}
                                />
                                <Box><Text color="status-critical"><ErrorMessage res={data?.updateContact}/></Text></Box>
                            </Box>
                        }
                    />
                </Box>
        </Layer>
    )
}

function ErrorMessage({res}: { res: { __typename: "Forbidden", message: string } | { __typename: "ServerError", message: string } | any}) {
    switch (res?.__typename) {
        case "Forbidden":
            return <>Доступ запрещен</>
        case "ServerError":
            return <>Ошибка сервера</>
    }

    return null
}