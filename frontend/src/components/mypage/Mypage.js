import React,{useState, useEffect} from 'react';
import {Container, Col,Row, Image} from 'react-bootstrap';
import {useRecoilValue} from 'recoil';
import {currentUser} from '../atoms';
import {getAvatar} from '../../lib/api/mypage';
import MypageRight from './MypageRight';
import MypageLeft from './MypageLeft';


const Mypage = () =>{
    const [avatar, setAvatar] = useState(null);
    const user = useRecoilValue(currentUser);
    
    useEffect(()=>{
        getAvatar(user).then((res)=>setAvatar(res.data.avatar_url));
    }, [user])

    return (
        <Container className="quiz__container">
            <Row>
                <Col sm={4}>
                <MypageLeft avatar={avatar}/>
                </Col>
                <Col sm={8}>
                <MypageRight/>
                </Col>
            </Row>
        </Container>
    )
}

export default Mypage;