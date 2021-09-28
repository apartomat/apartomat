import React, { useState } from "react"
import { useParams } from "react-router-dom"

import { useAuthContext } from "../common/context/auth/useAuthContext"
import { useWorkspace, WorkspaceUsersResult, WorkspaceProjectsListResult, WorkspaceProject } from "./useWorkspace"
import { useCreateProject } from "./useCreateProject"

import { Main, Box, Header, Heading, Text,
    Avatar, List, Button, Tip, Paragraph, Spinner, SpinnerExtendedProps,
    Layer, Form, FormField, TextInput } from "grommet"
import AnchorLink from "../common/AnchorLink"
import UserAvatar from "./UserAvatar";

interface RouteParams {
    id: string
};

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

export function Workspace () {
    const { user } = useAuthContext()
    let { id } = useParams<RouteParams>()
    const { data, loading, error } = useWorkspace(parseInt(id, 10))

    const [ showCreateProjectLayer, setShowCreateProjectLayer ] = useState(false)

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
                <Paragraph>Can't get workspace: {error}</Paragraph>
            </Main>
        );
    }

    switch (data?.workspace.__typename) {
        case "Workspace":
            const { workspace } = data;
            return (
                <Main>
                    <Header background="white" margin={{top:"large", horizontal:"large", bottom:"medium"}}>
                        <Box>
                            <Text size="xlarge" weight="bold" color="brand">apartomat</Text>
                        </Box>
                        <Box><UserAvatar user={user} className="header-user" /></Box>
                    </Header>

                    <Box margin={{horizontal: "large"}}>
                        <Box direction="row" margin={{vertical: "medium"}} justify="between">
                            <Heading level={2} margin="none">{workspace.name}</Heading>
                            <Box>
                                <Button color="brand" label="Новый проект" onClick={() => setShowCreateProjectLayer(true)} />
                            </Box>
                        </Box>

                        <Projects projects={workspace.projects.list} />

                        <Box margin={{top: "xlarge"}}>
                            <WorkspaceUsers users={workspace.users} />
                        </Box>
                    </Box>

                    {showCreateProjectLayer && <CreateProject workspaceId={workspace.id} />}
                </Main>
            );
        case "NotFound":
            return (
                <Main pad="large">
                    <Heading level={2}>Ошибка</Heading>
                    <Box>
                        <Text>Workspace not found</Text>
                    </Box>
                </Main>
            );
        case "Forbidden":
            return (
                <Main pad="large">
                    <Heading level={2}>Ошибка</Heading>
                    <Paragraph>Access is denied</Paragraph>
                </Main>
            );
        default:
            return (
                <Main pad="large">
                    <Heading>Ошибка</Heading>
                    <Paragraph>Unknown error</Paragraph>
                </Main>
            );
    }
}

function WorkspaceUsers({ users }: {users: WorkspaceUsersResult}) {
    switch (users.__typename) {
        case "WorkspaceUsers":
            return (
                <Box direction="row">
                    {users.items.map(user => {
                        return (
                            <Tip content={user.profile.email}>
                                <Avatar
                                    key={user.id}
                                    src={user.profile.gravatar?.url}
                                    size="medium"
                                    background="light-1"
                                    border={{ color: 'white', size: 'small' }}
                                >{user.profile.abbr}</Avatar>
                            </Tip>
                        )
                    })}
                </Box>
            )
        default:
            return null
    }
}

function Projects({ projects }: { projects: WorkspaceProjectsListResult }) {
    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            return (
                <List
                    primaryKey="name"
                    data={projects.items}
                    pad={{vertical:"small"}}
                    margin={{vertical: "medium"}}>
                    {(pr: WorkspaceProject) => (
                        <Box direction="row" justify="between">
                            <AnchorLink to={`/p/${pr.id}`}>{pr.name}</AnchorLink>
                            <Text>{pr.period}</Text>
                        </Box>
                    )}
                </List>
            )
        default:
            return <div>n/a</div>
    }
}

function CreateProject({ workspaceId }: { workspaceId: number }) {
    const [ title, setTitle ] = useState("")
    const [ create, { loading, error } ] = useCreateProject()

    const handleSubmit = async (event: React.FormEvent) => {
        event.preventDefault()
        create({ workspaceId, title: title })
    }

    const handleChangeTitle = (event: React.FormEvent<HTMLInputElement>) => {
        setTitle(event.currentTarget.value)
    }

    if (loading) {
        return (
            <div>Createing project...</div>
        )
    }

    return (
        <Layer>
            <Box pad="medium" gap="medium" width="medium">
                <Box direction="row" justify="between"align="center">
                    <Heading level={3} margin="none">Новый проект</Heading>
                    {/* <Button icon={ <FormClose/> } onClick={() => setShow(false)}/> */}
                </Box>
                <Form onSubmit={handleSubmit}>
                    {error ? <p>{error.message}</p> : null}
                    <FormField>
                        <TextInput onChange={handleChangeTitle} value={title} />
                    </FormField>
                    <Box direction="row">
                        <Button type="submit" primary label="Создать" />
                    </Box>
                </Form>
            </Box>
        </Layer>
    )
}

export default Workspace;