import React from 'react'

import {
  MainContainer,
  Item,
  IconContainer,
  OptionsContainer
} from './Navbar.styled';

const Navbar = () => {


  //testing, library, extractor, pedigree
  const [selected, setSelected] = React.useState('');

  return (
    <MainContainer>
      <IconContainer>
        <img src="/icons/egg.png"></img>
      </IconContainer>
      <OptionsContainer>
        <Item selected={selected == 'testing'}>
          <p>Testing</p>
        </Item>
        <Item selected={selected == 'library'}>
          <p>Library</p>
        </Item>
        <Item selected={selected == 'extractor'}>
          <p>Extractor</p>
        </Item>
        <Item selected={selected == 'pedigree'}>
          <p>Pedigree</p>
        </Item>
      </OptionsContainer>
    </MainContainer>
  )
}

export default Navbar;
