import { QueryResult, useApolloClient } from "@apollo/client"
import { ProjectPageScreenQuery, ProjectPageScreenQueryVariables, useProjectPageScreenQuery } from "api/graphql"

export function useProjectPage(id: string): QueryResult<ProjectPageScreenQuery, ProjectPageScreenQueryVariables> {
    const client = useApolloClient()

    const result = useProjectPageScreenQuery({
        client,
        errorPolicy: "all",
        variables: { id },
    })

    return { ...result }
}

export type { ProjectPageScreenProjectFragment as ProjectPage } from "api/graphql"
export type { ProjectPageScreenHouseFragment as House } from "api/graphql"
export type { ProjectPageScreenAlbumFragment as Album } from "api/graphql"
export type { ProjectPageScreenVisualizationFragment as Visualization } from "api/graphql"
