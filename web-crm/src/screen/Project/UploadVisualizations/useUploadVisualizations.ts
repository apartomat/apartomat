import { FetchResult, useApolloClient } from "@apollo/client"
import { UploadVisualizationsMutation, useUploadVisualizationsMutation, UploadVisualizationsMutationResult } from "api/types.d"

export function useUploadVisualizations(): [
    ({ projectId, files, roomId}: { projectId: string, files: File[], roomId?: string}) => Promise<FetchResult< UploadVisualizationsMutation>>,
    UploadVisualizationsMutationResult
] {
    const client = useApolloClient(); 

    const [ upload, result ] = useUploadVisualizationsMutation({ client, errorPolicy: "all" });

    return [({ projectId, files, roomId  }) => upload({
        variables: { projectId, files: files, roomId }
    }), result];
}

export default useUploadVisualizations;