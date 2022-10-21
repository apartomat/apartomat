import React from "react"

import { WorkspaceScreenProjectFragment, WorkspaceScreenArchiveProjectsFragment } from "../useWorkspace"

import { Box, Heading, List, Text } from "grommet"
import AnchorLink from "common/AnchorLink"

export default function Archive({ projects: { done: projects } }: { projects: WorkspaceScreenArchiveProjectsFragment }) {
    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            if (projects.items.length === 0) {
                return null
            }

            return (
                <Box margin={{bottom: "medium"}}>
                    <Box direction="row" margin={{vertical: "medium"}} justify="between">
                        <Heading level={3} margin="none">Архив</Heading>
                    </Box>
                    <List
                        primaryKey="name"
                        data={projects.items}
                        pad={{vertical:"small"}}
                        margin={{vertical: "medium"}}>
                        {(project: WorkspaceScreenProjectFragment) => (
                            <Box direction="row" justify="between">
                                <AnchorLink to={`/p/${project.id}`}>{project.name}</AnchorLink>
                                <Text>{project.period}</Text>
                            </Box>
                        )}
                    </List>
                </Box>
            )
        default:
            return null
    }
}