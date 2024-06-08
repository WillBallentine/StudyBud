using System;
using StudyBud.Data.Interfaces;
using StudyBud.Business.Interfaces;

namespace StudyBud.Business
{
	public class SyllabusBLL : ISyllabusBLL
	{
		private ISyllabusDAL _syllabusDal;

		public SyllabusBLL(ISyllabusDAL syllabusDal)
		{
			_syllabusDal = syllabusDal;
		}

		public void ProcessSyllabus()
		{
			//process syllabus and call DAL
		}
	}
}

