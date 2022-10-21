import React, { useCallback } from "react"

import { WorkspaceScreenProjectFragment, WorkspaceScreenCurrentProjectsFragment } from "../useWorkspace"

import { Box, List, Text } from "grommet"
import { FormClose } from "grommet-icons"
import AnchorLink from "common/AnchorLink"

export default function Projects({ projects: { current: projects } }: { projects: WorkspaceScreenCurrentProjectsFragment }) {
    const error = useCallback(() => {
            switch (projects.__typename) {
                case "Forbidden":
                    return "Доступ запрещен"
                case "ServerError":
                    return "Ошибка сервера"
            }
    
            return ""
    }, [ projects ])

    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            return (
                <List
                    primaryKey="name"
                    data={projects.items}
                    pad={{vertical:"small"}}
                    margin={{vertical: "medium"}}
                >
                    {({ id, name, period }: WorkspaceScreenProjectFragment) => (
                        <Box direction="row" justify="between">
                            <AnchorLink to={`/p/${id}`}>{name}</AnchorLink>
                            <Text>{period}</Text>
                        </Box>
                    )}
                </List>
            )
        default:
            return (
                <Box
                    pad="small"
                    round="small"
                    direction="row"
                    gap="small"
                    align="center"
                    background={{ color: "status-critical", opacity: "weak"}}
                    width={{ max: "medium" }}
                >
                    <Box border={{ color: "status-critical", size: "small"}} round="large">
                        <FormClose color="status-critical" size="medium"/>
                    </Box>
                    <Text weight="bold" size="medium">{error()}</Text>
                </Box>
            )
    }
}