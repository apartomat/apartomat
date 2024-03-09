import React, { useCallback } from "react"

import { WorkspaceScreenCurrentProjectsFragment } from "../useWorkspace"

import { Box, List, Text } from "grommet"
import { FormClose } from "grommet-icons"
import AnchorLink from "common/AnchorLink"

export default function Projects({
    projects: { current: projects },
}: {
    projects: WorkspaceScreenCurrentProjectsFragment
}) {
    const error = useCallback(() => {
        switch (projects.__typename) {
            case "Forbidden":
                return "Доступ запрещен"
            case "ServerError":
                return "Ошибка сервера"
        }

        return ""
    }, [projects])

    switch (projects.__typename) {
        case "WorkspaceProjectsList":
            return (
                <List
                    pad={{ vertical: "small" }}
                    margin={{ vertical: "medium" }}
                    data={projects.items}
                    itemKey="id"
                    primaryKey={({ id, name }) => (
                        <AnchorLink key={id} to={`/p/${id}`}>
                            {name}
                        </AnchorLink>
                    )}
                    secondaryKey={({ id, period }) => <Text key={`period-${id}`}>{period}</Text>}
                />
            )
        default:
            return (
                <Box
                    pad="small"
                    round="small"
                    direction="row"
                    gap="small"
                    align="center"
                    background={{ color: "status-critical", opacity: "weak" }}
                    width={{ max: "medium" }}
                >
                    <Box border={{ color: "status-critical", size: "small" }} round="large">
                        <FormClose color="status-critical" size="medium" />
                    </Box>
                    <Text weight="bold" size="medium">
                        {error()}
                    </Text>
                </Box>
            )
    }
}
