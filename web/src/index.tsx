import React, { StrictMode } from "react";
import ReactDOM from "react-dom";

import "./index.css";
import App from "./App";

import { ApolloClient, InMemoryCache, createHttpLink, ApolloProvider } from "@apollo/client";
import { setContext } from '@apollo/client/link/context';

const authLink = setContext((_, { headers }) => {
  const token = JSON.parse(localStorage.getItem("token") || `""`);
  return {
    headers: {
      ...headers,
      authorization: token ? `Bearer ${token}` : "",
    }
  }
});

const link = createHttpLink({ uri: "http://localhost:8010/graphql" });

const apolloClient = new ApolloClient({ cache: new InMemoryCache(), link: authLink.concat(link) });

ReactDOM.render(
  <StrictMode>
    <ApolloProvider client={apolloClient}>
      <App />
    </ApolloProvider>
  </StrictMode>,
  document.getElementById('root')
);