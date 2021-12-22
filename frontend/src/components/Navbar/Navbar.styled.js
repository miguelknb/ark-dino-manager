import styled from 'styled-components';
import media from 'styled-media-query';


export const MainContainer = styled.div`
  display: flex;
  flex-direction: row;
  padding: .5rem 1.5rem;
  height: 2.5rem;
  background-color: #1f1f1f;
  justify-content: flex-start;
`;

export const OptionsContainer = styled.div`
  display: flex;
  flex-direction: row;
  width: 40%;
  justify-content: space-evenly;

  ${media.lessThan("1100px")`
    width: 80%;
  `};
`

export const Item = styled.div`
  display: flex;
  flex-direction: column;
  align-items: center;
  justify-content: center;
  
  p {
    font-family: 'Roboto', sans-serif;
    font-size: 1.2rem;
    color: #ffff;
  }
`;

export const IconContainer = styled.div`
  display: flex;
  cursor: pointer;

  img {
    height: 2rem;
    width: 2rem;
  }
`