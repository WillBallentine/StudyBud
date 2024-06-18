using System;
using StudyBud.Models;

namespace StudyBud.Data.Interfaces
{
	public interface ISyllabusDAL
	{
		//returning a string so I can return the SyllabusID created by the entry
		string AddUserSyllabusContent(string userId, byte[] syllabusContent);
	}
}

