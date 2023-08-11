import { FetchResult, useApolloClient } from "@apollo/client"

import {
    MakeProjectPublicMutation, MakeProjectPublicMutationResult, useMakeProjectPublicMutation,
    MakeProjectNotPublicMutation, MakeProjectNotPublicMutationResult, useMakeProjectNotPublicMutation
} from "api/graphql"

export function useMakeProjectPublic(projectId: string): [
    () => Promise<FetchResult<MakeProjectPublicMutation>>,
    MakeProjectPublicMutationResult
] {
    const client = useApolloClient()
    const [ run, result ] = useMakeProjectPublicMutation({ client, errorPolicy: "all" })

    return [
        () => run({ variables: { projectId } }),
        result,
    ]
}

export function useMakeProjectNotPublic(projectId: string): [
    () => Promise<FetchResult<MakeProjectNotPublicMutation>>,
    MakeProjectNotPublicMutationResult
] {
    const client = useApolloClient()
    const [ run, result ] = useMakeProjectNotPublicMutation({ client, errorPolicy: "all" })

    return [
        () => run({ variables: { projectId } }),
        result,
    ]
}
