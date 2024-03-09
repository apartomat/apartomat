import React from "react"

import { WorkspaceScreenArchiveProjectsFragment } from "../useWorkspace"

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
                        pad={{vertical:"small"}}
                        margin={{vertical: "medium"}}
                        data={projects.items}
                        itemKey="id"
                        primaryKey={({ id, name }) => (
                            <AnchorLink key={id} to={`/p/${id}`} >{name}</AnchorLink>
                        )}
                        secondaryKey={({ id, period }) => (
                            <Text key={`period-${id}`}>{period}</Text>
                        )}
                    />
                </Box>
            )
        default:
            return null
    }
}
