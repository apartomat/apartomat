import { HashRouter } from "react-router-dom"
import {ApolloClient, ApolloProvider, createHttpLink, DefaultOptions, InMemoryCache} from "@apollo/client";

import { Grommet } from "grommet"

import { ProjectPage } from "pages"

const theme = {
    global: {
        font: {
            family: "Roboto",
            size: "18px",
            height: "20px",
        },
    },
    carousel: {
        animation: {
            duration: 400,
        },
        icons: {
            color: "brand",
        },
        disabled: {
            icons: {
                color: "grey",
            },
        },
    },
}

const link = createHttpLink({ uri: import.meta.env.VITE_APARTOMAT_API_URL })

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

export function App() {
    return (
        <Grommet theme={theme}>
            <ApolloProvider client={apolloClient}>
                <HashRouter>
                    <ProjectPage id={id(window.location.pathname)}/>
                </HashRouter>
            </ApolloProvider>
        </Grommet>
    )
}

function id(path) {
    return path.replace(/^\//, '').replace(/\/$/, '')

}
