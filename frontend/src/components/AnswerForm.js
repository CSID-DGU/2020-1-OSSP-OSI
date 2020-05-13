import React, {Fragment} from 'react';
import {Form} from 'react-bootstrap';
import ChoicesList from './ChoicesList';

const AnswerForm = (props) =>{
    const {addChoices, onChange, quiz, onRemoveChoice} = props;
    if(props.quiz.type === "mul_choices"){
        return (
            <Fragment>
            <Form.Label>Answer - {quiz.type}</Form.Label>
            <Form.Row>
            <ChoicesList choices={quiz.choices} addChoices={addChoices} onChange={onChange} quiz={quiz} onRemoveChoice={onRemoveChoice}/>
            </Form.Row>
            </Fragment>
            );
    }
    return (
        <Form.Group>
            <Form.Label>Answer {props.quiz.type}</Form.Label>
            <Form.Control name="answer" index={quiz.id} value={quiz.answer} onChange={(e)=>onChange(e)}></Form.Control>
        </Form.Group>
    )
}

export default AnswerForm;