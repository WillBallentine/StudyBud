namespace StudyBud.Models;

public class Assignment
{
    public string AssignmentId { get; set; }

    public string Name { get; set; }

    public string Type { get; set; } //this might be a good place for an enum?

    public string Description { get; set; }

    public decimal PercentOfGrade { get; set; }



}


