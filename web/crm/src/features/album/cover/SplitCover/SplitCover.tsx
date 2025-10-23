import styled from "styled-components"

export function SplitCover({
    title,
    subtitle,
    qrcodeSrc,
    city,
    year,
    imgSrc,
    scale,
    logoText,
}: {
    title: string
    subtitle?: string | null
    qrcodeSrc?: string | null
    city?: string | null
    year?: number | null
    imgSrc: string
    scale: number
    logoText: string
}) {
    return (
        <Page $scale={scale}>
            <link href="https://fonts.googleapis.com/css2?family=Arsenal:wght@400&display=swap" rel="stylesheet" />
            <link href="https://fonts.googleapis.com/css2?family=Oswald:wght@400&display=swap" rel="stylesheet" />
            <Layout>
                <Left>
                    <Header></Header>
                    <Center>
                        <Title>{title}</Title>
                        <Subtitle>{subtitle}</Subtitle>
                        {qrcodeSrc && (
                            <QrCode>
                                <img src={qrcodeSrc} />
                            </QrCode>
                        )}
                    </Center>
                    <Footer>
                        <City>{city}</City>
                        <Year>•{year}•</Year>
                    </Footer>
                </Left>
                <Right>
                    <img src={imgSrc} />
                    <Logo>{logoText}</Logo>
                </Right>
            </Layout>
        </Page>
    )
}

const Page = styled.div<{ $scale?: number }>`
    font:
        400 1em Arsenal,
        sans-serif;
    background-color: white;

    transform: scale(${(props) => props.$scale || 1.0});
    transform-origin: left top;
    width: 297mm;
    height: 211mm;
`

const Layout = styled.div`
    width: 297mm;
    height: 211mm;
    display: flex;
`

const Left = styled.div`
    flex: 0 0 50%;
    display: flex;
    flex-direction: column;
    justify-content: space-between;
`

const Header = styled.div``

const Center = styled.div`
    display: flex;
    flex-direction: column;
    align-items: center;
`

const Title = styled.div`
    display: flex;
    justify-content: left;
    font-size: 42px;
    font-weight: 100;
    text-transform: uppercase;
`

const Subtitle = styled.div`
    display: flex;
    font-size: 14px;
    text-align: center;
    max-width: 50%;
`

const Footer = styled.div`
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    margin-bottom: 20px;
`

const City = styled.div`
    font-size: 12px;
`

const Year = styled.div`
    font-size: 12px;
`

const Right = styled.div`
    flex: 0 0 50%;
    position: relative;

    img {
        width: 100%;
        height: 100%;
        object-fit: cover;
    }
`
const Logo = styled.div`
    display: flex;
    width: fit-content;

    color: #a67034;
    border: #a67034 1px solid;
    border-radius: 7px;
    padding: 3px 4px;
    font:
        400 1em Oswald,
        sans-serif;
    font-size: 14px;
    line-height: 14px;

    ${Right} & {
        position: absolute;
        top: 20px;
        left: 50%;
        transform: translateX(-50%);
    }
`
const QrCode = styled.div`
    position: relative;
    top: 200px;

    img {
        width: 48px;
        height: 48px;
        border: #9b9c9b 1px solid;
        border-radius: 3px;
        padding: 3px;
    }
`
