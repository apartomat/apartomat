import styled from "styled-components"

const Wrapper = styled.div`
    display: flex;
    flex-direction: row;
    justify-content: center;
    margin-top: 50px;
`

const Page = styled.div`
    // background-color: #eee;
    width: 297mm;
    height: 211mm;
    zoom: 100%;
    // border: #f00 1px solid;
    font:
        400 1em Oswald,
        sans-serif;
    box-shadow:
        rgba(50, 50, 93, 0.25) 0px 50px 100px -20px,
        rgba(0, 0, 0, 0.3) 0px 30px 60px -30px;
`

const Layout = styled.div`
    // background-color: #cfe;
    height: 100%;
    width: 100%;
    display: flex;
`

const Left = styled.div`
    flex: 50%;
    // background-color: #8f0;
    display: flex;
    flex-direction: column;
    // padding: 1cm;
    justify-content: space-between;
`

const Logo = styled.div`
    display: flex;
    justify-content: left;
    font-size: 26pt;
    color: rgb(82, 42, 40);
    margin-top: 1cm;
    margin-left: 1cm;
`

const Project = styled.div`
    height: 500px;
    margin-left: 1cm;
`

const Title = styled.div`
    // background-color: cyan;
    display: flex;
    justify-content: left;
    font-size: 64pt;
    font-weight: 100;
`

const Subtitle = styled.div`
    // background-color: magenta;
    display: flex;
    justify-content: left;
    margin-left: 10px;
    font-size: 18pt;
    margin-top: 50px;
`

const City = styled.div`
    // background-color: blue;
    display: flex;
    justify-content: left;
    font-size: 14pt;
    margin-left: 10px;
`

const Right = styled.div`
    flex: 50%;
    // background-color: #0dc;
`

export function AlbumCover() {
    return (
        <Wrapper>
            <Page>
                <link
                    href="https://fonts.googleapis.com/css2?family=Oswald:wght@200,400&display=swap"
                    rel="stylesheet"
                />
                <Layout>
                    <Left>
                        <Logo>PUHOVA</Logo>
                        <Project>
                            <Title>TESLA</Title>
                            <City>Новосибирск</City>
                            <Subtitle>Дизайн-проект интерьера квартиры 66м²</Subtitle>
                        </Project>
                    </Left>
                    <Right>
                        <img
                            src="http://localhost:9000/apartomat/p/kIX2JEMeEZHKDZulcPMgg/tDgk6a8H0u8f41aUQuUPi.jpg"
                            style={{ objectFit: "cover", width: "100%", height: "100%" }}
                        />
                    </Right>
                </Layout>
            </Page>
        </Wrapper>
    )
}
