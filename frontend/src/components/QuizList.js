import React, {Fragment} from "react";
import QuizForm from "./QuizForm";
import AnswerForm from "./AnswerForm";


const QuizList = (props) => {
    const { quizzes, onRemove, onTypeChange, addChoices, onChange} = props;

    const QuizLists = quizzes.map((quiz, index) => 
        <QuizForm key={quiz.id} onRemove={onRemove} onTypeChange={onTypeChange} AnswerForm={AnswerForm} addChoices={addChoices} quiz={quiz} onChange={onChange} />
        );

    return (
        <Fragment>
            {QuizLists}
        </Fragment>
    );
};

export default QuizList;