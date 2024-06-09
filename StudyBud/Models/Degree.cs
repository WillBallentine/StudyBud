namespace StudyBud.Models;

public class Degree
{
	public string DegreeId { get; set; }

	public string DegreeType { get; set; }

	public School IssuingSchool { get; set; }

	public float GPA { get; set; }

	public DateTime YearStarted { get; set; }

	public DateTime YearFinished { get; set; }

	public bool Graduated { get; set; }
}

