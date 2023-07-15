export type TypeInterface = {
    title: string;
    description: string;
    buttonText: string;
} 

export type TypeClasses = {
    classtitle: string,
    subtitle: string,
    textbutton: string,
}

export type TypeOur = {
    data: TypeInterface,
    classes?: TypeClasses | null,
}