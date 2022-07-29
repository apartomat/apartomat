import { FetchResult, useApolloClient } from "@apollo/client"
import { useCreateProjectMutation, CreateProjectMutation, CreateProjectMutationResult, Project, Forbidden, ServerError } from "api/types.d"

export type CreateProjectFn = ({workspaceId, title, startAt, endAt }:{workspaceId: string, title: string, startAt?: Date, endAt?: Date })
    => Promise<FetchResult<CreateProjectMutation>>;

export enum State {
    INITIAL = "INITIAL",
    CREATING = "CREATING",
    DONE = "DONE",
    FAILED = "FAILED"
}

export type Result =
    Initial |
    Creating |
    Done |
    Failed


type Initial = {
    state: State.INITIAL
}

type Creating = {
    state: State.CREATING
}

type Done = {
    state: State.DONE
    project: Pick<Project, 'id' | 'title'>
}

type Failed = {
    state: State.FAILED
    error: Error | (
        { __typename: 'Forbidden' }
        & Pick<Forbidden, 'message'>
        ) | (
        { __typename: 'ServerError' }
        & Pick<ServerError, 'message'>
        )
}

export function useCreateProject(): [
    CreateProjectFn,
    CreateProjectMutationResult,
    Result
] {
    const client = useApolloClient(); 
    const [create, result] = useCreateProjectMutation({ client, errorPolicy: 'all' });

    const state = new Proxy<Result>({ state: State.INITIAL }, {
        get: (target: Result, name: string) => {
            switch (name) {
                case "state":
                    if (!result.called) {
                        return State.INITIAL
                    } else if (result.called && result.loading) {
                        return State.CREATING
                    } else if (result.called && !result.loading &&
                        (result.error || (result.data && result.data?.createProject.__typename !== "ProjectCreated"))
                    ) {
                        return State.FAILED
                    } else if (result.called && !result.loading && result.data && result.data?.createProject.__typename === "ProjectCreated") {
                        return State.DONE
                    }

                    return State.INITIAL

                case "project":
                    if (result.data?.createProject.__typename === "ProjectCreated") {
                        return result.data.createProject
                    }

                    return undefined

                case "error":
                    if (result.data && result.data?.createProject.__typename !== "ProjectCreated") {
                        return result.data.createProject
                    }
                    
                    return result.error

                default:
                    return undefined
            }
        }
    })

    return [({ workspaceId, title, startAt, endAt }) => create({ variables: { input: { workspaceId, title, startAt, endAt }} }), result, state as Result]
}

export default useCreateProject;