using System;
using StudyBud.Models;

namespace StudyBud.Business.Interfaces
{
	public interface ISyllabusBLL
	{
		bool ProcessSyllabus(MemoryStream syllabus, string userId);
	}
}

