import React, { useEffect } from "react";
import {Route, Redirect, RouteProps } from "react-router-dom";

import { Main, Box, Text, SpinnerExtendedProps, Spinner, Heading, Paragraph } from "grommet"

import useAuthContext, { UserContextStatus } from "../useAuthContext"

const Loading = (props: SpinnerExtendedProps) => {
    return (
        <Spinner
            border={[
                { side: 'all', color: 'background-contrast', size: 'medium' },
                { side: 'right', color: 'brand', size: 'medium' },
                { side: 'top', color: 'brand', size: 'medium' },
                { side: 'left', color: 'brand', size: 'medium' },
            ]}
            {...props}
        />
    )
}


function PrivateRoute({ children, ...rest }: RouteProps) {
    const { user, check, error } = useAuthContext();

    useEffect(() => {
        if (user.status === UserContextStatus.UNDEFINED) {
            check();
        }
    }, [user, check]);

    if (error !== undefined) {
        return (
            <Main pad="large">
                <Heading level={2}>Ошибка</Heading>
                <Paragraph>Не удалось получить профиль пользователя</Paragraph>
            </Main>
        );
    }

    return (
        <Route {...rest} render={({ location }) => {
            switch (user.status) {
                case UserContextStatus.UNDEFINED:
                case UserContextStatus.CHEKING:
                    return (
                        <Main pad="large">
                            <Box direction="row" gap="small" align="center">
                                <Loading message="Авторизация..."/>
                                <Text>Авторизация...</Text>
                            </Box>
                        </Main>
                    );
                case UserContextStatus.SERVER_ERROR:
                    return (
                        <Main pad="large">
                            <Heading level={2}>Ошибка</Heading>
                            <Paragraph>Не удалось получить профиль пользователя</Paragraph>
                        </Main>
                    );
                case UserContextStatus.LOGGED:
                    return children;
                default:
                    return <Redirect to={{ pathname: "/login", state: { from : location }}}/>;
            }
        }} />
    );
}

export default PrivateRoute;