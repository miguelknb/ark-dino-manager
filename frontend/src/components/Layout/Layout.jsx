import React from 'react'
import { createGlobalStyle } from 'styled-components';

import Footer from '../Footer/Footer';
import Navbar from '../Navbar/Navbar';

import {
  MainContainer,
  Content
} from './Layout.styled';


const GlobalStyle = createGlobalStyle`
  body {
    margin: 0;
    padding: 0;
    font-family: 'Roboto', sans-serif;
  }
`;

export const Layout = ({children}) => {
  return (
    <>
      <GlobalStyle/>
      <MainContainer>
        <Navbar/>
        <Content>
          {children}
        </Content>
        <Footer/>
      </MainContainer>
    </>
  )
}

export default Layout;