import React from 'react'

import {
  MainContainer
} from './Layout.styled';

export const Layout = ({children}) => {
  return (
    <MainContainer>
      <Navbar/>
      {children}
      <Footer/>
    </MainContainer>
  )
}

export default Layout;