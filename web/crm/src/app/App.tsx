import { Grommet } from "grommet"

import { ApolloClient, InMemoryCache, ApolloProvider, split, DefaultOptions } from "@apollo/client"
import { setContext } from "@apollo/client/link/context"
import { createUploadLink } from "apollo-upload-client"
import { getMainDefinition } from "@apollo/client/utilities"

import { BrowserRouter, Routes, Route } from "react-router-dom"

import { AuthProvider, NotificationsProvider } from "./providers/"
import { AuthRequired, RedirectToDefaultWorkspace } from "./routes"

import { AcceptInvite, Album, AlbumCover, Confirm, Login, Logout, Project, Visualizations, Workspace } from "pages"
import { GraphQLWsLink } from "@apollo/client/link/subscriptions"
import { createClient } from "graphql-ws"

const theme = {
    global: {
        font: {
            family: "Roboto",
            size: "18px",
            height: "20px",
        },
    },
}

const authLink = setContext((req, { headers }) => {
    const token = JSON.parse(localStorage.getItem("token") || `""`)

    return {
        headers: {
            ...headers,
            authorization: token ? `Bearer ${token}` : "",
        },
    }
})

const httpLink = createUploadLink({ uri: import.meta.env.VITE_APARTOMAT_API_URL })

const wsLink = new GraphQLWsLink(
    createClient({
        url: import.meta.env.VITE_APARTOMAT_API_URL_WS,
        connectionParams: () => {
            return {
                Authorization: JSON.parse(localStorage.getItem("token") || `""`),
            }
        },
    })
)

const link = split(
    ({ query }) => {
        const definition = getMainDefinition(query)
        return definition.kind === "OperationDefinition" && definition.operation === "subscription"
    },
    wsLink,
    authLink.concat(httpLink)
)

const cache = new InMemoryCache()

const defaultOptions: DefaultOptions = {
    watchQuery: {
        fetchPolicy: "no-cache",
        errorPolicy: "all",
    },
    query: {
        fetchPolicy: "no-cache",
        errorPolicy: "all",
    },
}

const apolloClient = new ApolloClient({ cache, link, defaultOptions })

const messages = {
    messages: {
        input: {
            readOnlyCopy: {
                prompt: "В буфер обмена",
                validation: "️Скопировано",
            },
        },
    },
} as Parameters<typeof Grommet>[0]["messages"]

function App() {
    return (
        <Grommet theme={theme} messages={messages}>
            <ApolloProvider client={apolloClient}>
                <AuthProvider>
                    <NotificationsProvider>
                        <BrowserRouter>
                            <Routes>
                                <Route path="/login" element={<Login />} />
                                <Route path="/logout" element={<Logout />} />
                                <Route path="/confirm" element={<Confirm />} />
                                <Route path="/accept-invite" element={<AcceptInvite />} />
                                <Route element={<AuthRequired />}>
                                    <Route path="/" element={<RedirectToDefaultWorkspace />} />
                                    <Route path="/:id" element={<Workspace />} />
                                    <Route path="/p/:id" element={<Project />} />
                                    <Route path="/vis/:id" element={<Visualizations />} />
                                    <Route path="/album/:id" element={<Album />} />
                                    <Route path="/album/:id/cover" element={<AlbumCover />} />
                                </Route>
                            </Routes>
                        </BrowserRouter>
                    </NotificationsProvider>
                </AuthProvider>
            </ApolloProvider>
        </Grommet>
    )
}

export default App
