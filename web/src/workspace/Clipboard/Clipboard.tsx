import React, { useState, useEffect }from 'react';
import './Clipboard.css';

import CircularProgress from '@material-ui/core/CircularProgress';

function useLoading(): [boolean, (fn: () => void) => void] {
    const [ isLoading, setLoading ] = useState(false);
  
    const wrapper = async (fn: () => void) => {
        setLoading(true);
        await fn();
        setLoading(false);
    };
  
    return [ isLoading, wrapper ];
}

export type Product = {
    name: string
    description: string
    image: string
};

async function findProductOnPage(url: string): Promise<Product | null> {

    const q = {
        query: `query ($url: String!) {
            shoppingList {
              findProductOnPage(url: $url) {
                  name
                  description
                  image
                }
            }
          }`,
        variables: {
            url
        }
    }

    const resp = await fetch(
        'http://localhost:8010/graphql',
        {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json;charset=UTF-8'
            },
            redirect: 'follow',
            body: JSON.stringify(q)
        }
    );

    const json: { data: { shoppingList: { findProductOnPage: Product }} } = await resp.json();

    return json.data.shoppingList.findProductOnPage;
}

async function setContentFromClipboard(
    content: string,
    setContent: (value: React.SetStateAction<string>) => void,
    setProduct: (value: React.SetStateAction<Product | null>) => void,
    loading: (fn: () => void) => void
  ) {
  
    try {
        const txt = await navigator.clipboard.readText();
        if (content !== txt) {
            loading(async () => {
                const product = await findProductOnPage(txt);
                setProduct(product);
                setContent(txt);
            });
        }
    } catch (e) {
        //
    }
}

function useClipboardContent(ms: number): [string, Product | null, Boolean] {
    const [ content, setContent ] = useState("");
    const [ isLoading, loading ] = useLoading();

    const [ product, setProduct ] = useState<Product | null>(null);
  
    useEffect(() => {
      const t = setInterval(() => {
        setContentFromClipboard(content, setContent, setProduct, loading); 
      }, ms);
  
      return () => clearInterval(t);
    }, [content, ms, loading]);
  
    return [content, product, isLoading];
}

function ProductAdd({ prod, url, onAdd }: {prod: Product, url: string, onAdd: (item: Product) => void}) {
    return (
        <span>
            <a href={url}>{ prod.name }</a>&nbsp;
            <button onClick={() => onAdd(prod)}>+</button>
        </span>
    );
}

export function Clipboard({ onAdd }: { onAdd: (item: Product) => void }) {
    const [content, product, isLoading] = useClipboardContent(300);
  
    return (
        <div className="clipboard">
            Clipboard: { isLoading && <CircularProgress/> } { !isLoading && product && <ProductAdd prod={product} url={content} onAdd={onAdd} /> }
        </div>
    );
}

export function ProductList({ list }: {list: Product[]}) {
    return (
        <ul>
            { list.map(item => (<ProductListItem key={item.name} item={item}/>)) }
        </ul>
    );
}

function ProductListItem({ item }: { item: Product }) {
    return (
        <li>{ item.name } <img width="100px" height="100px" src={item.image} alt="" /></li>
    );
}