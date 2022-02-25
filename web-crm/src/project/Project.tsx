import React, { ChangeEvent, Dispatch, SetStateAction, useEffect, useState, useRef } from "react"
import { useParams } from "react-router-dom"

import { Main, Box, Header, Heading, Text,
    Paragraph, Spinner, SpinnerExtendedProps, Image,
FileInput, Button, Layer, Form, FormField, Drop, MaskedInput, Tip } from "grommet"
import { FormClose, StatusGood, Add } from "grommet-icons"

import AnchorLink from "../common/AnchorLink"

import UserAvatar from "./UserAvatar"

import { useProject, ProjectFiles, ProjectContacts, ProjectHouses, HouseRooms } from "./useProject"
import { useUploadProjectFile, ProjectFileType } from "./useUploadProjectFile"
import { useAddContact, ContactType } from "./useAddContact"

import { useAuthContext } from "../common/context/auth/useAuthContext"

interface RouteParams {
    id: string
};

export function Project () {
    let { id } = useParams<RouteParams>()

    const { user } = useAuthContext()
    const { data, loading, error } = useProject(parseInt(id, 10))

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

    const [showAddContact, setShowAddContact] = useState(false);

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
            const { project } = data?.screen.projectScreen;

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
                    <Box margin={{vertical: "medium"}}>
                        <Heading level={2} margin="none">{project.title}</Heading>
                    </Box>
                    <Box direction="row" justify="between" wrap>
                        <Box width={{min: "35%"}}>
                            <Box margin="none">
                                <Heading level={4} margin={{ bottom: "xsmall"}}>Сроки проекта</Heading>
                                <ProjectDates startAt={project.startAt} endAt={project.endAt}/>
                            </Box>
                            <Contacts contacts={project.contacts} showAddContact={setShowAddContact}/>
                        </Box>
                        <Box width={{min: "35%"}}>
                            <House houses={project.houses}/>
                            <Rooms houses={project.houses}/>
                        </Box>
                    </Box>

                    <Visualizations files={project.files} showUploadFiles={setShowUploadFiles}/>

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

                    <AddSomething showUploadFiles={setShowUploadFiles}/>

                    {showUploadFiles ?
                        <UploadFiles
                            projectId={project.id}
                            type={ProjectFileType.Visualization}
                            setShow={setShowUploadFiles}
                            onUploadComplete={({message}) => notify({ message })}
                        /> : null}

                    {showAddContact ?
                        <AddContact
                            projectId={project.id.toString()}
                            setShow={setShowAddContact}
                            onAdded={({ message }) => notify({ message })}
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

function Contacts({ contacts, showAddContact }: { contacts: ProjectContacts, showAddContact: Dispatch<SetStateAction<boolean>> }) {
    switch (contacts.list.__typename) {
        case "ProjectContactsList":
            return (
                <Box margin={{top: "small"}}>
                    <Heading level={4} margin={{ bottom: "xsmall"}}>
                        {contacts.list.items.length === 1 ? "Заказчик" : "Заказчики"}
                    </Heading>
                    <Box direction="row" gap="small" wrap>
                        {[...contacts.list.items.map((contact) => {
                            return (
                                <Button key={contact.id} primary color="light-2" label={contact.fullName} style={{whiteSpace: "nowrap"}}/>
                            )
                        }), <Button icon={<Add/>} label="Добавить" onClick={() => showAddContact(true) }/>]}
                    </Box>
                </Box>
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

function LocaleDate ({ children }: { children: string }) {
    const date = new Date(children)

    return (
        <>{date.toLocaleDateString('ru-RU')}</>
    )
}

function ProjectDates ({ startAt, endAt }: {startAt?: string, endAt?: string}) {
    const [ state, setState ] = useState<string[]>(
        [startAt, endAt].reduce((acc, item) => { if (item) { acc.push(item); } return acc; }, [] as string[])
    )
    const [ label, setLabel ] = useState(<>не определены</>)

    useEffect(() => {
        if (state[0] && state[1] && state[0] !== state[1]) {
            setLabel(<>{new Date(state[0]).toLocaleDateString("ru-RU")}&nbsp;&mdash;&nbsp;{new Date(state[1]).toLocaleDateString("ru-RU")}</>)
            return
        }

        if (state[0]) {
            setLabel(<>{new Date(state[0]).toLocaleDateString("ru-RU")}</>)
            return
        }

        setLabel(<>не определены</>)
    }, [state])

    return (
        <Box direction="row">
            <Button primary color="light-2" label={label}/>
        </Box>
    )
}

function AddSomething ({ showUploadFiles }: { showUploadFiles: Dispatch<SetStateAction<boolean>> }) {
    const [show, setShow] = useState(false)

    const targetRef = useRef<HTMLDivElement>(null)

    return (
        <Box align="center" margin={{top:"xlarge"}}>

            <Box ref={targetRef}>
                <Button primary label="Добавить" onClick={() => setShow(true)} size="large" />
            </Box>

            {show && targetRef.current && (
                <Drop
                    elevation="small"
                    round="small"
                    align={{bottom: "top"}}
                    target={targetRef.current}
                    margin={{bottom:"small"}}
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
            
            <Box margin={{top: "medium"}}>
                <Text>Добавьте визуализации, планы, исходники, создайте альбом и спецификацию.</Text>
            </Box>
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
        projectId: number,
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
    { projectId, setShow, onAdded }:
    {
        projectId: string,
        setShow: Dispatch<SetStateAction<boolean>>,
        onAdded: ({ message }: { message: string }) => void
    }) {

    const [ value, setValue ] = useState({fullName: "", phone: "", email: "", instagram: ""})

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
                setShow(false)
                onAdded({ message: "Контакт добавлен" })
        }
    }, [ data, setShow, onAdded ])

    return (
        <Layer
            onClickOutside={() => setShow(false)}
            onEsc={() => setShow(false)}
        >
            <Form
                validate="submit"
                value={value}
                onChange={val => setValue(val)}
                onSubmit={handleSubmit}
                messages={{required: 'обязательное поле'}}
            >
                {error && <Box><Text>{error.message}</Text></Box>}

                <Box pad="medium" gap="medium" width={{min: "500px"}}>
                    <Box direction="row"justify="between"align="center">
                        <Heading level={3} margin="none">Добавить контакт</Heading>
                        <Button icon={ <FormClose/> } onClick={() => setShow(false)}/>
                    </Box>
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
                    <Box direction="row" justify="between" margin={{top: "small"}}>
                        <Box><Text color="status-critical"><ErrorMessage res={data?.addContact}/></Text></Box>
                        <Button
                            type="submit"
                            primary
                            label={loading ? 'Сохранение...' : 'Сохранить' }
                            disabled={loading}
                        />
                    </Box>
                </Box>
            </Form>
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