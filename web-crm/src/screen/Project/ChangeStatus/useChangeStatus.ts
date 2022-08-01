import { FetchResult, useApolloClient } from "@apollo/client"
import { useChangeProjectStatusMutation, ChangeProjectStatusMutation, ProjectStatus, ChangeProjectStatusMutationResult } from "api/types.d"

export function useChangeStatus(): [
    (projectId: string, status: ProjectStatus) => Promise<FetchResult<ChangeProjectStatusMutation>>,
    ChangeProjectStatusMutationResult
] {
    const client = useApolloClient()
    const [ change, result ] = useChangeProjectStatusMutation({ client, errorPolicy: 'all' })

    return [
        (projectId: string, status: ProjectStatus) => change({ variables: { projectId, status } }),
        result,
    ]
}
