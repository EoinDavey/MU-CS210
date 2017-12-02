main=interact((\j->let n=foldl(\x (w:ws)->case w of "INSERT"->x++[ws!!0];_|x==[]->[];_->tail x)[] j in case length n of 0->"empty\n";_->head n++['\n']).map words.lines)
